package forms

import "assay/dao"

type LoginForm struct {
	Username string `json:"username" binding:"required_if=Type 0"`
	Password string `json:"password" binding:"required_if=Type 0"`
	Phone    string `json:"phone" binding:"required_if=Type 1"`
	Code     string `json:"code" binding:"required_if=Type 1"`
	Type     int    `json:"type" binding:"omitempty,oneof=0 1" comment:"0-密码登录 1-短信登录"`
	PlatForm string `json:"platform" binding:"required,oneof=assay" comment:"平台类型"`
}

type Login struct {
	Token string `json:"token"`
	dao.User
}
