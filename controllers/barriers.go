package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type BarriersController struct{}

func BarriersRegister(router *gin.RouterGroup) {
	controller := &BarriersController{}

	// 开闸放车
	router.POST("open", controller.open)
}

func (*BarriersController) open(c *gin.Context) {
	params := &forms.BarriersOpenForm{}

	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	barriersService.Open(c, params)
}
