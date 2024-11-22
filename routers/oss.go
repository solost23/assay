package routers

import (
	"assay/controllers"
	"github.com/gin-gonic/gin"
)

func InitOSSRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("oss")
	{
		controllers.OSSRegister(routerGroup)
	}
}
