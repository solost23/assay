package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("users")
	{
		controllers.UserRegister(routerGroup)
	}
}
