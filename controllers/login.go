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
}

func (*LoginController) login(c *gin.Context) {
	params := &forms.LoginForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	loginService.Login(c, params)
}
