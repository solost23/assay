package initialize

import (
	"assay/infra/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	WebConfigPath = "configs/config.yml"
)

func initConfig() {
	v := viper.New()
	v.SetConfigFile(WebConfigPath)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic(err)
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panic(err)
	}
}
