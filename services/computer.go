package services

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ComputerService struct{}

func (*ComputerService) Status(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer conn.Close()

	// 接收数据
	go func(conn *websocket.Conn) {
		for {
			record := &forms.ComputerStatus{}
			if err := conn.ReadJSON(record); err != nil {
				zap.S().Info("client disconnected")
				return
			}
			fmt.Println("record: ", record)
		}
	}(conn)

	// 发送数据
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			record := forms.ComputerStatus{
				Type: 1,
				Data: "hello, world!",
			}
			if err := conn.WriteJSON(record); err != nil {
				zap.S().Info("client disconnected")
				return
			}
		}
	}
}
