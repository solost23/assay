package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type TextController struct{}

func TextRegister(router *gin.RouterGroup) {
	controller := &TextController{}

	// 文本屏配置获取
	router.GET("config", controller.config)
	// 文本屏设置
	router.POST("setting", controller.setting)
}

func (*TextController) config(c *gin.Context) {
	textConfig := global.ServerConfig.Text
	response.Success(c, forms.TextConfig{
		Width:  textConfig.Width,
		Height: textConfig.Height,
	})
}

func (*TextController) setting(c *gin.Context) {
	params := &forms.TextSettingForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	textService.Setting(c, params)
}
