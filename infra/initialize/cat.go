package initialize

import (
	"assay/infra/global"
	"github.com/tarm/serial"
	"go.uber.org/zap"
	"time"
)

func initCat() {
	catConfig := global.ServerConfig.Cat

	s, err := serial.OpenPort(&serial.Config{
		Name:        catConfig.Name,
		Baud:        catConfig.Baud,
		ReadTimeout: time.Duration(catConfig.ReadTimeout) * time.Millisecond,
		Size:        catConfig.Size,
		Parity:      serial.Parity(catConfig.Parity),
		StopBits:    serial.StopBits(catConfig.StopBits),
	})
	if err != nil {
		zap.S().Panic("failed to connect cat: ", err)
	}
	global.Cat = s
}
