package forms

type TextConfig struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type TextSettingForm struct {
	Name   string  `json:"name"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type       int     `json:"type" comment: 0-动态 -1-静态`
	LineHeight float64 `json:"lineHeight"`
	Speed      int     `json:"speed"`
	Html       string  `json:"html"`
	Left       int     `json:"left"`
	Top        int     `json:"top"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
}
