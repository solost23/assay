package servants

import (
	"assay/forms"
	"assay/infra/global"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

const (
	Static = iota - 1
	Dynaic
)

const (
	timeout = 300 * time.Millisecond
)

func NewXiXunPlayRequest(name string, blocks []forms.Block) *XiXunPlayRequest {
	layers := make([]Layer, 0, len(blocks))
	for i := 0; i != len(blocks); i++ {
		// 默认是动态
		name := "SingleText"
		class := "SingleLineText"
		if blocks[i].Type == Static {
			name = "MultiText"
			class = "MultiLineText"
		}
		source := Source{
			ID:                  uuid.NewString(),
			Name:                name,
			Type:                class,
			LineHeight:          blocks[i].LineHeight,
			Speed:               blocks[i].Speed,
			HTML:                blocks[i].Html,
			PlayTime:            0,
			TimeSpan:            10,
			Left:                blocks[i].Left,
			Top:                 blocks[i].Top,
			Width:               blocks[i].Width,
			Height:              blocks[i].Height,
			EntryEffect:         "None",
			ExitEffect:          "None",
			EntryEffectTimeSpan: 0,
			ExitEffectTimeSpan:  0,
		}
		layer := Layer{
			Repeat:  false,
			Sources: []Source{source},
		}
		layers = append(layers, layer)
	}
	request := &XiXunPlayRequest{
		Type: "commandXixunPlayer",
		ID:   uuid.NewString(),
		Command: Command{
			Type: "PlayXixunTask",
			ID:   uuid.New().String(),
			Task: Task{
				ID:     uuid.NewString(),
				Name:   fmt.Sprintf("%s_task", name),
				Insert: false,
				Items: []Item{
					{
						ID: uuid.NewString(),
						Program: Program{
							ID:        uuid.NewString(),
							TotalSize: 0,
							Name:      name,
							Width:     global.ServerConfig.Text.Width,
							Height:    global.ServerConfig.Text.Height,
							Layers:    layers,
						},
						RepeatTimes: 1,
						Schedules: []Schedule{
							{
								DateType:   "All",
								TimeType:   "Range",
								StartTime:  "00:00",
								EndTime:    "23:59",
								FilterType: "None",
							},
						},
					},
				},
			},
		},
	}

	request.client = resty.New()
	request.client.SetTimeout(timeout)

	return request
}

func (r *XiXunPlayRequest) Send() (*Response, error) {
	resp := &Response{}

	_, err := r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(r).
		SetResult(resp).
		Post(global.ServerConfig.Text.Url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type XiXunPlayRequest struct {
	Type    string        `json:"type"`
	ID      string        `json:"_id"`
	Command Command       `json:"command"`
	client  *resty.Client `json:"-"`
}

type Command struct {
	Type string `json:"_type"`
	ID   string `json:"id"`
	Task Task   `json:"task"`
}

type Task struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Insert bool   `json:"insert"`
	Items  []Item `json:"items"`
}

type Item struct {
	ID          string     `json:"_id"`
	Program     Program    `json:"_program"`
	RepeatTimes int        `json:"repeatTimes"`
	Schedules   []Schedule `json:"schedules"`
}

type Program struct {
	ID        string  `json:"_id"`
	TotalSize int     `json:"totalSize"`
	Name      string  `json:"name"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Layers    []Layer `json:"layers"`
}

type Layer struct {
	Repeat  bool     `json:"repeat"`
	Sources []Source `json:"sources"`
}

type Source struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Type                string  `json:"_type"`
	LineHeight          float64 `json:"lineHeight"`
	Speed               int     `json:"speed"`
	HTML                string  `json:"html"`
	PlayTime            int     `json:"playTime"`
	TimeSpan            int     `json:"timeSpan"`
	Left                int     `json:"left"`
	Top                 int     `json:"top"`
	Width               int     `json:"width"`
	Height              int     `json:"height"`
	EntryEffect         string  `json:"entryEffect"`
	ExitEffect          string  `json:"exitEffect"`
	EntryEffectTimeSpan int     `json:"entryEffectTimeSpan"`
	ExitEffectTimeSpan  int     `json:"exitEffectTimeSpan"`
}

type Schedule struct {
	DateType   string `json:"dateType"`
	TimeType   string `json:"timeType"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	FilterType string `json:"filterType"`
}

type Response struct {
	Type      string `json:"_type"`
	ID        string `json:"_id"`
	Timestamp int64  `json:"timestamp"`
	DeviceId  string `json:"deviceId"`
}
