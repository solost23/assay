package main

import (
	"assay/dao"
	"assay/infra/global"
	"assay/infra/initialize"
	"assay/infra/shutdown"
	"assay/routers"
	"assay/services"
	"assay/services/servants"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	if err := global.DB.AutoMigrate(&dao.Role{}, &dao.User{}, &dao.Device{}, &dao.Alarm{}); err != nil {
		zap.S().Panic("failed to migrate database: ", err)
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
