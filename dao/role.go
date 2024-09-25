package dao

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;type:varchar(100);comment: 名称"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(100);comment: 昵称"`
}
