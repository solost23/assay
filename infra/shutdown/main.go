package shutdown

import (
	"assay/infra/global"

	"go.uber.org/zap"
)

func Destroy() {
	zap.L().Info("start destroy resources.")
	// mysql
	sqlDB, _ := global.DB.DB()
	if err := sqlDB.Close(); err != nil {
		zap.S().Error(err)
	}
	// redis
	if err := global.RDB.Close(); err != nil {
		zap.S().Error(err)
	}
	// mqtt
	mqttConfig := global.ServerConfig.Mqtt
	global.Mqtt.Disconnect(mqttConfig.Quiesce)
	// cat
	if err := global.Cat.Close(); err != nil {
		zap.S().Error(err)
	}
	zap.L().Info("success destroy resources.")
}
