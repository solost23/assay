package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitTextScreenRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("text")
	{
		controllers.TextRegister(routerGroup)
	}
}
