package initialize

import (
	"assay/infra/global"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(register func() *gin.Engine) {
	serverConfig := global.ServerConfig

	gin.SetMode(serverConfig.Mode)
	r := register()
	go func() {
		if err := r.Run(fmt.Sprintf(":%d", serverConfig.Port)); err != nil {
			zap.S().Panic("server run failed", err)
		}
	}()
}
