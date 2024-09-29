package controllers

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/infra/util"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func UserRegister(router *gin.RouterGroup) {
	controller := &UserController{}

	// 用户添加
	router.POST("", controller.insert)
	// 用户删除
	router.DELETE(":id", controller.delete)
	// 用户修改
	router.PUT(":id", controller.update)
	// 用户列表
	router.GET("", controller.list)
}

func (*UserController) insert(c *gin.Context) {
	params := &forms.UserInsertForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	userService.Insert(c, params)
}

func (*UserController) delete(c *gin.Context) {
	uIdForm := &constant.UIdForm{}
	if err := util.GetValidUriParams(c, uIdForm); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	userService.Delete(c, uIdForm.Id)
}

func (*UserController) update(c *gin.Context) {
	uIdForm := &constant.UIdForm{}
	if err := util.GetValidUriParams(c, uIdForm); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	params := &forms.UserInsertForm{}
	if err := util.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}

	userService.Update(c, uIdForm.Id, params)
}

func (*UserController) list(c *gin.Context) {
	params := &forms.UserListForm{}
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

	userService.List(c, params)
}
