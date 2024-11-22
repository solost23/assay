package servants

import (
	"assay/constants"
	"assay/infra/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SubscriptionDeviceStatus(callback mqtt.MessageHandler) error {
	// TODO: 订阅主题暂时采用设备增删改主题
	return Subscription(constants.AssayDeviceTopic, callback)
}

func SubscriptionAlarm(callback mqtt.MessageHandler) error {
	return Subscription(constants.AssayAlarmTopic, callback)
}

func Subscription(topic string, callback mqtt.MessageHandler) error {
	client := global.Mqtt

	if token := client.Subscribe(topic, 1, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
