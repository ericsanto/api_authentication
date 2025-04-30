package user

import (
	"errors"
	"net/http"

	"github.com/ericsanto/api_authentication/internal/customerrors"
	"github.com/ericsanto/api_authentication/pkg/validate"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) PostUser(c *gin.Context) {

	var requestUser UserRequest

	switch {
	case !validate.ValidateJSONBody(&requestUser, c):
		return

	case !validate.FieldsBodyIsVallid(c, requestUser):
		return
	}

	userResponse, err := uc.userService.CreateUserService(&requestUser)
	if err != nil {

		switch {
		case customerrors.IsDuplicateKeyUniqueConstraint(err):
			c.JSON(http.StatusConflict, customerrors.Conflict(err.Error()))
			return
		case !customerrors.IsDuplicateKeyUniqueConstraint(err):
			c.JSON(http.StatusInternalServerError, customerrors.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusCreated, userResponse)

}

func (uc *UserController) Login(c *gin.Context) {

	var UserRequestLogin UserRequestLogin

	if !validate.ValidateJSONBody(&UserRequestLogin, c) {
		c.JSON(http.StatusBadRequest, customerrors.BadRequest())
		return
	}

	if !validate.FieldsBodyIsVallid(c, UserRequestLogin) {
		return
	}

	userResponse, err := uc.userService.Login(UserRequestLogin)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNotFound):
			c.JSON(http.StatusNotFound, customerrors.NotFound(err.Error()))
			return

		case errors.Is(err, customerrors.ErrPasswordIncorrect):
			c.JSON(http.StatusUnauthorized, customerrors.Unauthorized(err.Error()))
			return

		default:
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	tokenString, err := uc.userService.GenerateToken(userResponse.ID, userResponse.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
