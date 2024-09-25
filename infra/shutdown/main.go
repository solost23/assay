package shutdown

import (
	"assay/infra/global"

	"go.uber.org/zap"
)

func Destroy() {
	zap.L().Info("start destroy resources.")
	sqlDB, _ := global.DB.DB()
	if err := sqlDB.Close(); err != nil {
		zap.S().Error(err)
	}
	zap.L().Info("success destroy resources.")
}
