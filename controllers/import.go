package controllers

import (
	"assay/infra/constant"
	"assay/infra/response"
	"github.com/gin-gonic/gin"
)

type ImportController struct{}

func ImportRegister(router *gin.RouterGroup) {
	controller := &ImportController{}

	// 任务导入模板
	router.GET("template/tasks", controller.templateTasks)
	// 任务导入
	router.POST("upload/tasks", controller.uploadTasks)
}

func (*ImportController) templateTasks(c *gin.Context) {
	importService.TemplateTasks(c)
}

func (*ImportController) uploadTasks(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	importService.UploadTasks(c, file)
}
