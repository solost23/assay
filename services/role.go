package services

import (
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleService struct{}

func (*RoleService) Insert(c *gin.Context, params *forms.RoleInsertForm) {
	db := global.DB

	// 检查角色是否存在
	_, err := dao.GWhereFirstSelect[dao.Role](db, "id", "name = ?", params.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if err == nil {
		response.Error(c, constant.BadRequestCode, errors.New("角色已存在，不允许重复创建"))
		return
	}
	if len(params.RouterAuth) != 0 {

	}
	if len(params.ButtonAuth) != 0 {

	}

	err = dao.GInsert(db, &dao.Role{
		Name:     params.Name,
		Nickname: params.Nickname,
	})
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}
