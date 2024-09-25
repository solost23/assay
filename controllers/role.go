package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func RoleRegister(router *gin.RouterGroup) {
	controller := &RoleController{}

	// 角色添加
	router.POST("", controller.insert)
}

func (*RoleController) insert(c *gin.Context) {
	params := &forms.RoleInsertForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	roleService.Insert(c, params)
}
