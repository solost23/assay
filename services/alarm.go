package services

import (
	"assay/dao"
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
)

type AlarmService struct{}

func (*AlarmService) List(c *gin.Context, params *forms.AlarmListForm) {
	query := []string{"1 = ?"}
	args := []any{1}
	if params.StartTime != "" {
		startTime, err := now.ParseInLocation(time.Local, params.StartTime)
		if err != nil {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		query = append(query, "created_at >= ?")
		args = append(args, startTime)
	}
	if params.EndTime != "" {
		endTime, err := now.ParseInLocation(time.Local, params.EndTime)
		if err != nil {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		query = append(query, "created_at <= ?")
		args = append(args, endTime)
	}

	db := global.DB
	if params.Keyword != "" {
		sqlDevices, err := dao.GWhereAllSelectOrder[dao.Device](db, "id", "id DESC", "name LIKE ?", params.Keyword)
		if err != nil {
			response.Error(c, constant.InternalServerErrorCode, err)
			return
		}
		deviceIds := make([]uint, 0, len(sqlDevices))
		for i := 0; i != len(sqlDevices); i++ {
			deviceIds = append(deviceIds, sqlDevices[i].ID)
		}
		query = append(query, "device_id IN (?)")
		args = append(args, deviceIds)
	}

	sqlAlarms, total, pages, err := dao.GPaginateOrder[dao.Alarm](db, &dao.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, "id desc", strings.Join(query, " AND "), args...)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	deviceIds := make([]uint, 0, len(sqlAlarms))
	for i := 0; i != len(sqlAlarms); i++ {
		deviceIds = append(deviceIds, sqlAlarms[i].DeviceId)
	}

	sqlDevices, err := dao.GWhereAllSelectOrder[dao.Device](db, "id, name", "id DESC", "id IN (?)", deviceIds)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	deviceIdMap := make(map[uint]string, len(sqlDevices))
	for i := 0; i != len(sqlDevices); i++ {
		deviceIdMap[sqlDevices[i].ID] = sqlDevices[i].Name
	}

	records := make([]forms.AlarmListRecord, 0, len(sqlAlarms))
	for i := 0; i != len(sqlAlarms); i++ {
		records = append(records, forms.AlarmListRecord{
			ID:         sqlAlarms[i].ID,
			Level:      sqlAlarms[i].Level,
			DeviceId:   sqlAlarms[i].DeviceId,
			DeviceName: deviceIdMap[sqlAlarms[i].DeviceId],
			FaultType:  sqlAlarms[i].FaultType,
			EndTime:    sqlAlarms[i].EndTime.Format(time.DateTime),
			Interval:   uint(sqlAlarms[i].EndTime.Sub(sqlAlarms[i].CreatedAt).Minutes()),
			CreatedAt:  sqlAlarms[i].CreatedAt.Format(time.DateTime),
		})
	}

	response.Success(c, &forms.AlarmList{
		PageList: constant.PageList{
			Current: params.Page,
			Pages:   pages,
			Size:    params.Size,
			Total:   total,
		},
		Records: records,
	})
}

func (*AlarmService) InsertAlarmTask(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("当前话题是%s, 信息是%s", message.Topic(), string(message.Payload()))
	// TODO: 解析 json 数据存入数据库
}
