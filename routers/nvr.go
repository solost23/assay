package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitNvrRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("nvr")
	{
		controllers.NvrRegister(routerGroup)
	}
}
