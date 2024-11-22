package initialize

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}
