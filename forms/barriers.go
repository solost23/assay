package forms

type BarriersOpenForm struct {
	ParkCode         string `json:"parkCode"`
	PlateNo          string `json:"plateNo"`
	CardNo           string `json:"cardNo"`
	PassTime         string `json:"passTime"`
	LaneCode         string `json:"laneCode"`
	VehicleType      int    `json:"vehicleType"`
	PlateColor       int    `json:"plateColor"`
	GateName         string `json:"gateName"`
	LaneName         string `json:"laneName"`
	UniqueNo         string `json:"uniqueNo"`
	PicFilePath      string `json:"picFilePath"`
	PicPlateFilePath string `json:"picPlateFilePath"`
}
