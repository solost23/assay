package services

import (
	"assay/forms"
	"assay/infra/constant"
	"assay/infra/response"
	"assay/services/servants"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BarriersService struct{}

func (*BarriersService) Open(c *gin.Context, params *forms.BarriersOpenForm) {
	// 检查是否是允许通行，此处先都允许通行
	zap.S().Infof("recv: %+#v from plateNo %s", params, params.PlateNo)
	// 发送 mqtt
	publishBarriersForm := servants.PublishBarriersForm{
		CarCode:   params.PlateNo,
		CarStatus: 1,
		TaskId:    1,
	}
	if err := servants.PublishBarriers(publishBarriersForm); err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"errDesc": "success",
	})
}
