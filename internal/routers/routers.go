package routers

import (
	"github.com/ericsanto/api_authentication/internal/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	router := gin.Default()

	user.UserRouter(router)

	return router

}
