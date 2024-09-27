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

	router.GET("", controller.list)
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
