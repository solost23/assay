package initialize

import (
	"assay/infra/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func initMqtt() {
	mqttConfig := global.ServerConfig.Mqtt

	opts := mqtt.NewClientOptions().
		AddBroker(mqttConfig.Addr).
		SetClientID(uuid.New().String()).
		SetUsername(mqttConfig.Username).
		SetPassword(mqttConfig.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		zap.S().Panic("failed to connect mqtt: ", token.Error())
	}
	global.Mqtt = client
}
