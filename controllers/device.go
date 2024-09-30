package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type DeviceController struct{}

func DeviceRegister(router *gin.RouterGroup) {
	controller := &DeviceController{}

	// 设备添加
	router.POST("", controller.insert)
	// 同步设备状态
	router.GET("status", controller.status)
}

func (*DeviceController) insert(c *gin.Context) {
	params := &forms.DeviceInsertForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	deviceService.Insert(c, params)
}

func (*DeviceController) status(c *gin.Context) {
	deviceService.Status(c)
}
