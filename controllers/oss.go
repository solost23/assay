package controllers

import (
	"assay/infra/constant"
	"assay/infra/response"
	"github.com/gin-gonic/gin"
)

type OSSController struct{}

func OSSRegister(router *gin.RouterGroup) {
	controller := OSSController{}

	// 静态文件存储
	router.POST("static", controller.staticUpload)
	// 动态文件存储
	router.POST("dynamic", controller.dynamicUpload)
}

func (*OSSController) staticUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	OSSService.StaticUpload(c, file)
}

func (*OSSController) dynamicUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	OSSService.DynamicUpload(c, file)
}
