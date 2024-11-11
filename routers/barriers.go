package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitBarriersRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("barriers")
	{
		controllers.BarriersRegister(routerGroup)
	}
}
