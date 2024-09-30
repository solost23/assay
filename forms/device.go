package forms

type DeviceInsertForm struct {
	Name    string `json:"name" binding:"required"`
	Host    string `json:"host" binding:"required"`
	Port    int    `json:"port" binding:"required"`
	Protocl int    `json:"protocl" binding:"omitempty,oneof=-1 0 1"`
}

type DeviceStatus struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
