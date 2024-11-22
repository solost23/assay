package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitDeviceRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("devices")
	{
		controllers.DeviceRegister(routerGroup)
	}

}
