package dao

import "gorm.io/gorm"

const (
	DeviceProtocolTcp = iota - 1
	DeviceProtocolModuleBus
	DeviceProtocolDDS
)

const (
	DeviceStatusFree = iota - 1
	DeviceStatusWork
	DeviceStatusClose
)

type Device struct {
	gorm.Model
	Name     string `json:"name" gorm:"column:name;type:varchar(500);comment: 设备名称"`
	Host     string `json:"host" gorm:"column:host;type:varchar(100);comment: 主机地址"`
	Port     int    `json:"port" gorm:"column:port;type:int;comment: 端口号"`
	Protocol int    `json:"protocol" gorm:"column:protocol;type:int;comment: 通信协议 -1-Tcp 0-ModuleBus 1-DDS"`
	Status   int    `json:"status" gorm:"column:status;type:tinyint;comment: 运行状态 -1-空闲 0-工作 1-关闭"`
}
