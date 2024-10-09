package dao

import (
	"time"

	"gorm.io/gorm"
)

const (
	AlarmLevelOne = iota - 1
	AlarmLevelTwo
	AlarmLevelThree
)

const (
	AlarmFaultTypeDisconnect = iota - 1
	AlarmFaultTypePowerOutage
	AlarmFaultTypeFault
)

// Alarm 警告信息表
type Alarm struct {
	gorm.Model
	TaskId    uint      `json:"taskId" gorm:"column:task_id;type:bigint unsigned;comment: 任务 ID"`
	Level     int       `json:"level" gorm:"column:level;type:tinyint;comment: 警告级别 -1-一级 0-二级 1-三级"`
	DeviceId  uint      `json:"deviceId" gorm:"column:device_id;type:bigint unsigned;comment: 设备 ID"`
	FaultType int       `json:"faultType" gorm:"column:fault_type;type:tinyint;comment: 故障类型 -1-断网 0-断电 1-故障"`
	EndTime   time.Time `json:"endTime" gorm:"column:end_time;type:time;comment: 消除警报时间"`
}
