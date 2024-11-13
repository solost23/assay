package services

import (
	"assay/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BarriersService struct{}

func (*BarriersService) Open(c *gin.Context, params *forms.BarriersOpenForm) {
	// 检查是否是允许通行，此处先都允许通行
	zap.S().Infof("recv: %+#v from plateNo %s", params, params.PlateNo)
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"errDesc": "success",
	})
}
