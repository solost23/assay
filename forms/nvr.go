package forms

type NvrChannel struct {
	Data []NvrChannelData `json:"data"`
}

type NvrChannelData struct {
	ChanName string `json:"chanName"`
	ChanNum  int    `json:"chanNum"`
	DeviceId string `json:"deviceId"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}
