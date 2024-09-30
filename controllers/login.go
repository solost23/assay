package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func LoginRegister(router *gin.RouterGroup) {
	controller := &LoginController{}

	// 登陆
	router.POST("", controller.login)
	// 短信猫发送随机4位验证码
	router.GET("code", controller.code)
}

func (*LoginController) login(c *gin.Context) {
	params := &forms.LoginForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	loginService.Login(c, params)
}

func (*LoginController) code(c *gin.Context) {
	params := &forms.LoginGetCodeForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	loginService.Code(c, params)
}
