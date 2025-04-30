package validate

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ericsanto/api_authentication/internal/customerrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSONBody(request interface{}, c *gin.Context) bool {

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, customerrors.BadRequest())
		return false
	}

	return true
}

func ValidateFieldErrors422UnprocessableEntity(entity interface{}) (map[string]string, error) {

	valid := validator.New()

	err := valid.Struct(entity)
	if err != nil {

		validatorErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)

		for _, fieldErrors := range validatorErrors {
			errorMessages[strings.ToLower(fieldErrors.Field())] = fmt.Sprintf("%s: %s", fieldErrors, fieldErrors.Tag())
		}

		return errorMessages, nil
	}

	return nil, err
}

func FieldsBodyIsVallid(c *gin.Context, request interface{}) bool {

	validate, err := ValidateFieldErrors422UnprocessableEntity(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, customerrors.InternalServerError())
		return false
	}

	if len(validate) > 0 {
		c.JSON(http.StatusUnprocessableEntity, customerrors.UnprocessableEntity(validate))
		return false
	}

	return true
}
