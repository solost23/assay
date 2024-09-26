package dao

import "gorm.io/gorm"

const (
	DeviceStatusFree = iota - 1
	DeviceStatusWork
	DeviceStatusClose
)

type Device struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name;type:varchar(500);comment: 设备名称"`
	Status int    `json:"status" gorm:"column:status;type:tinyint;comment: 运行状态 -1-空闲 0-工作 1-关闭"`
}
