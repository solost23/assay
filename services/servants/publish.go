package servants

import (
	"assay/constants"
	"assay/dao"
	"assay/infra/global"
	"encoding/json"
)

func PublishDevices() error {
	db := global.DB

	sqlDevices, err := dao.GWhereAllSelectOrder[dao.Device](db, "*", "id DESC", "1 = ?", 1)
	if err != nil {
		return err
	}
	return Publish(constants.AssayDeviceTopic, sqlDevices)
}

func Publish(topic string, data any) error {
	client := global.Mqtt

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if token := client.Publish(topic, 1, false, b); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
