package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitAlarmRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("alarms")
	{
		controllers.AlarmRegister(routerGroup)
	}
}
