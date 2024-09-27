package main

import (
	"assay/dao"
	"assay/infra/global"
	"assay/infra/initialize"
	"assay/infra/shutdown"
	"assay/routers"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	if err := global.DB.AutoMigrate(&dao.Role{}, &dao.User{}, &dao.Device{}, &dao.Alarm{}); err != nil {
		zap.S().Panic("failed to migrate database: ", err)
	}

	initialize.Run(routers.Register)
	defer shutdown.Destroy()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
