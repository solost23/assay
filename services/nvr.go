package services

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type NvrService struct{}

const (
	timeout = 3 * time.Second
)

func (*NvrService) Channel(c *gin.Context) {
	client := resty.New()
	client.SetTimeout(timeout)

	resp := &forms.NvrChannel{}
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(resp).
		Get(fmt.Sprintf("%s/api/hik/nvr/channel", global.ServerConfig.Hik.Url))
	if err != nil {
		response.Error(c, constant.BadRequestCode, err)
		return
	}
	response.Success(c, resp.Data)
}
