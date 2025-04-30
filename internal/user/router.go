package user

import (
	"github.com/ericsanto/api_authentication/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	userRepositoty := NewUserRepository(config.DB)
	userService := NewUserService(userRepositoty)
	userController := NewUserController(userService)

	userRouter := r.Group("/v1/auth")

	userRouter.POST("/", userController.PostUser)
	userRouter.POST("/login", userController.Login)

}
