package services

import (
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type DeviceService struct{}

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
