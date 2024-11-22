package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitComputerRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("computers")
	{
		controllers.ComputerRegister(routerGroup)
	}
}
