package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;type:varchar(100);comment: 用户名"`
	Password string `json:"password" gorm:"column:password;type:varchar(300);comment: 密码"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(100);comment: 昵称"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(20);comment: 手机号"`
	Email    string `json:"email" gorm:"column:email;type:varchar(100);comment: 邮箱"`
	RoleId   uint   `json:"roleId" gorm:"column:role_id;type:bigint unsigned;comment: 角色 ID"`
}
