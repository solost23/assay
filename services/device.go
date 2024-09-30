package services

import (
	"assay/constants"
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"encoding/json"
	"errors"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeviceService struct{}

func (*DeviceService) Insert(c *gin.Context, params *forms.DeviceInsertForm) {
	db := global.DB

	_, err := dao.GWhereFirstSelect[dao.Device](db, "id", "name = ? OR host = ? AND port = ?", params.Name, params.Host, params.Port)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if err == nil {
		response.Error(c, constant.BadRequestCode, errors.New("设备已存在，参数错误"))
		return
	}
	sqlDevice := &dao.Device{
		Host:     params.Host,
		Name:     params.Name,
		Port:     params.Port,
		Protocol: params.Protocl,
		Status:   dao.DeviceStatusFree,
	}
	if err := dao.GInsert(db, sqlDevice); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	// 发送到mqtt server
	mqttConfig := global.ServerConfig.Mqtt

	opts := mqtt.NewClientOptions().
		AddBroker(mqttConfig.Addr).
		SetClientID(uuid.New().String()).
		SetUsername(mqttConfig.Username).
		SetPassword(mqttConfig.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer client.Disconnect(mqttConfig.Quiesce)
	b, err := json.Marshal(sqlDevice)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if token := client.Publish(constants.AssayDeviceInsertTopic, 1, false, b); token.Wait() && token.Error() != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}

func (*DeviceService) Status(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			db := global.DB
			sqlDevices, err := dao.GWhereAllSelectOrder[dao.Device](db, "*", "id DESC", "1 = ?", 1)
			if err != nil {
				response.Error(c, constant.InternalServerErrorCode, err)
				return
			}
			records := make([]forms.DeviceStatus, 0, len(sqlDevices))
			for i := 0; i != len(sqlDevices); i++ {
				records = append(records, forms.DeviceStatus{
					ID:     sqlDevices[i].ID,
					Name:   sqlDevices[i].Name,
					Status: sqlDevices[i].Status,
				})
			}
			if err := conn.WriteJSON(records); err != nil {
				zap.S().Info("client disconnected")
				return
			}
		}
	}
}
