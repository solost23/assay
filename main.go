package main

import (
	"assay/dao"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/initialize"
	"assay/infra/shutdown"
	"assay/infra/util"
	"assay/routers"
	"assay/services"
	"assay/services/servants"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	db := global.DB
	if err := db.AutoMigrate(&dao.Role{}, &dao.User{}, &dao.Device{}, &dao.Alarm{}, &dao.OSSFile{}); err != nil {
		zap.S().Panic("failed to migrate database: ", err)
	}

	// 创建默认角色和用户
	sqlUserCnt, err := dao.GWhereCount[dao.User](db, "1 = ?", 1)
	if err != nil {
		zap.S().Panic("failed to query user cnt: ", err)
	}
	if sqlUserCnt == 0 {
		sqlRole := &dao.Role{
			Name:     "admin",
			Nickname: "admin",
		}
		if err := dao.GInsert(db, sqlRole); err != nil {
			zap.S().Panic("failed create default role: ", err)
		}
		sqlUser := &dao.User{
			Username: "admin",
			Password: util.NewMd5("123", constant.Secret),
			Nickname: "admin",
			// TODO: 后面补充
			Phone:  "",
			Email:  "",
			RoleId: sqlRole.ID,
		}
		if err := dao.GInsert(db, sqlUser); err != nil {
			zap.S().Panic("failed create default user: ", err)
		}
	}

	// 订阅任务
	// TODO: 此处考虑逻辑是否下放
	// 设备状态更新
	go func() {
		if err := servants.SubscriptionDeviceStatus((&services.DeviceService{}).UpdateStatusTask); err != nil {
			zap.S().Panic("subscription device status err: ", err)
		}
		zap.S().Info("subscription device status success")
	}()
	// 警报信息添加
	go func() {
		if err := servants.SubscriptionAlarm((&services.AlarmService{}).InsertAlarmTask); err != nil {
			zap.S().Panic("subscription alarm err: ", err)
		}
		zap.S().Info("subscription alarm success")
	}()
	initialize.Run(routers.Register)
	defer shutdown.Destroy()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
