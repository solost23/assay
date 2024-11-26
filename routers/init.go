package routers

import (
	"assay/middlewares"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	apiGroup := router.Group("api/assay")
	{
		InitLoginRouter(apiGroup)
		// 上位机状态同步
		InitComputerRouter(apiGroup)
		// 文件存储
		InitOSSRouter(apiGroup)
		// excel导入
		InitImportRouter(apiGroup)
		// 设备报警
		InitAlarmRouter(apiGroup)
		// 文本屏
		InitTextScreenRouter(apiGroup)
		// 道闸
		InitBarriersRouter(apiGroup)
		// NVR
		InitNvrRouter(apiGroup)
	}
	apiGroup.Use(
		middlewares.JWTAuth(),
	)
	{
		InitRoleRouter(apiGroup)
		InitUserRouter(apiGroup)
		// 设备状态同步
		InitDeviceRouter(apiGroup)
	}

	return router
}
