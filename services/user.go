package services

import (
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"assay/infra/util"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct{}

func (*UserService) Insert(c *gin.Context, params *forms.UserInsertForm) {
	db := global.DB

	// 检查角色是否存在
	_, err := dao.GWhereFirstSelect[dao.Role](db, "id", "id = ?", params.RoleId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.BadRequestCode, errors.New("角色不存在，参数错误"))
		return
	}

	// 检查当前用户是否存在
	_, err = dao.GWhereFirstSelect[dao.User](db, "id", "username = ? OR phone = ? OR email = ?", params.Username, params.Phone, params.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if err == nil {
		response.Error(c, constant.BadRequestCode, errors.New("用户已存在，不允许重复创建"))
		return
	}

	err = dao.GInsert(db, &dao.User{
		Username: params.Username,
		Password: util.NewMd5(params.Password, constant.Secret),
		Nickname: params.Nickname,
		Phone:    params.Phone,
		Email:    params.Email,
		RoleId:   params.RoleId,
	})
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}

func (*UserService) Delete(c *gin.Context, id uint) {
	db := global.DB
	if err := dao.GDelete[dao.User](db, "id = ?", id); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	response.Success(c, "success")
}

func (*UserService) Update(c *gin.Context, id uint, params *forms.UserInsertForm) {
	db := global.DB

	sqlUser, err := dao.GWhereFirstSelect[dao.User](db, "id", "id = ?", id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.BadRequestCode, errors.New("用户不存在，参数错误"))
		return
	}

	_, err = dao.GWhereFirstSelect[dao.Role](db, "id", "id = ?", params.RoleId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.BadRequestCode, errors.New("角色不存在，参数错误"))
		return
	}

	sqlUser.Username = params.Username
	sqlUser.Password = util.NewMd5(params.Password, constant.Secret)
	sqlUser.Nickname = params.Nickname
	sqlUser.Phone = params.Phone
	sqlUser.Email = params.Email
	sqlUser.RoleId = params.RoleId

	if err := dao.GSave[dao.User](db, sqlUser, "id = ?", id); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}

func (*UserService) List(c *gin.Context, params *forms.UserListForm) {
	db := global.DB

	query := []string{"1 = ?"}
	args := []any{1}
	if params.Username != "" {
		query = append(query, "username LIKE ?")
		args = append(args, "%"+params.Username+"%")
	}
	sqlUsers, total, pages, err := dao.GPaginateOrder[dao.User](db, &dao.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, "id DESC", strings.Join(query, " AND "), args...)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	roleIds := make([]uint, 0, len(sqlUsers))
	for i := 0; i != len(sqlUsers); i++ {
		roleIds = append(roleIds, sqlUsers[i].RoleId)
	}
	sqlRoles, err := dao.GWhereAllSelectOrder[dao.Role](db, "id, name", "id DESC", "id IN ?", roleIds)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	roleIdMap := make(map[uint]string, len(sqlRoles))
	for i := 0; i != len(sqlRoles); i++ {
		roleIdMap[sqlRoles[i].ID] = sqlRoles[i].Name
	}

	records := make([]forms.UserListRecord, 0, len(sqlUsers))
	for i := 0; i != len(sqlUsers); i++ {
		records = append(records, forms.UserListRecord{
			ID:        sqlUsers[i].ID,
			Username:  sqlUsers[i].Username,
			Nickname:  sqlUsers[i].Nickname,
			Phone:     sqlUsers[i].Phone,
			Email:     sqlUsers[i].Email,
			Role:      roleIdMap[sqlUsers[i].RoleId],
			UpdatedAt: sqlUsers[i].UpdatedAt.Format(time.DateTime),
		})
	}

	response.Success(c, &forms.UserList{
		PageList: constant.PageList{
			Current: params.Page,
			Pages:   pages,
			Size:    params.Size,
			Total:   total,
		},
		Records: records,
	})
}
