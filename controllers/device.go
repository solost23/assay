package controllers

import "github.com/gin-gonic/gin"

type DeviceController struct{}

func DeviceRegister(router *gin.RouterGroup) {
	controller := &DeviceController{}

	// 同步设备状态
	router.GET("status", controller.status)
}

func (*DeviceController) status(c *gin.Context) {
	deviceService.Status(c)
}
