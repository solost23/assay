package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type AlarmController struct{}

func AlarmRegister(router *gin.RouterGroup) {
	controller := &AlarmController{}

	// 报警信息保存
	router.POST("", controller.insert)
	// 报警列表
	router.GET("", controller.list)
}

func (*AlarmController) insert(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/xml")

	params := &forms.AlarmInsertForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	alarmService.Insert(c, params)
}

func (controller *AlarmController) list(c *gin.Context) {
	params := &forms.AlarmListForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 10
	}

	alarmService.List(c, params)
}
