package services

import (
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"assay/services/servants"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"

	"github.com/gin-gonic/gin"
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
		Protocol: params.Protocol,
		Status:   dao.DeviceStatusFree,
	}
	if err = dao.GInsert(db, sqlDevice); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	// TODO: 暂时不做成异步发送
	if err = servants.PublishDevices(); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, "success")
}

func (*DeviceService) Delete(c *gin.Context, id uint) {
	db := global.DB

	if err := dao.GDelete[dao.Device](db, "id = ?", id); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	if err := servants.PublishDevices(); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	response.Success(c, "success")
}

func (*DeviceService) Update(c *gin.Context, id uint, params *forms.DeviceInsertForm) {
	db := global.DB

	sqlDevice, err := dao.GWhereFirstSelect[dao.Device](db, "*", "id = ?", id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constant.BadRequestCode, errors.New("设备不存在，参数错误"))
		return
	}

	sqlDevice.Name = params.Name
	sqlDevice.Host = params.Host
	sqlDevice.Port = params.Port
	sqlDevice.Protocol = params.Protocol

	if err = dao.GSave[dao.Device](db, sqlDevice, "id = ?", id); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	if err = servants.PublishDevices(); err != nil {
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

func (*DeviceService) UpdateStatusTask(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("当前话题是%s, 信息是%s", message.Topic(), string(message.Payload()))
	// TODO: 解析 json 数据存入数据库
}
