package services

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/services/servants"

	"github.com/gin-gonic/gin"
)

type TextService struct{}

func (*TextService) Setting(c *gin.Context, params *forms.TextSettingForm) {
	if _, err := servants.NewXiXunPlayRequest(params.Name, params.Blocks).Send(); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	response.Success(c, "success")
}
