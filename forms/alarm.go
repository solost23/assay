package forms

import "assay/infra/constant"

type AlarmListForm struct {
	constant.PageForm
	StartTime string `form:"startTime"`
	EndTime   string `form:"endTime"`
	Keyword   string `form:"keyword"`
}

type AlarmList struct {
	constant.PageList
	Records []AlarmListRecord `json:"records"`
}

type AlarmListRecord struct {
	ID         uint   `json:"id"`
	Level      int    `json:"level"`
	DeviceId   uint   `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	FaultType  int    `json:"faultType"`
	EndTime    string `json:"endTime"`
	Interval   uint   `json:"interval"`
	CreatedAt  string `json:"createdAt"`
}
