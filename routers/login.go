package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitLoginRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("login")
	{
		controllers.LoginRegister(routerGroup)
	}
}
