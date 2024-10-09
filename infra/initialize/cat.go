package initialize

import (
	"assay/infra/global"
	"github.com/tarm/serial"
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
		// TODO: 暂时注释
		//zap.S().Panic("failed to connect cat: ", err)
	}
	global.Cat = s
}
