package services

import (
	"assay/infra/constant"
	"assay/infra/global"
	"assay/infra/response"
	"assay/infra/util"
	"assay/services/servants"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type OSSService struct{}

func (*OSSService) StaticUpload(c *gin.Context, file *multipart.FileHeader) {
	staticOSSConfig := global.ServerConfig.StaticOSS
	url, err := servants.UploadFile(0, staticOSSConfig.Bucket, file.Filename, file, servants.FileTypeImage)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, util.FulfillImageOSSPrefix(util.TrimDomainPrefix(url)))
}

func (*OSSService) DynamicUpload(c *gin.Context, file *multipart.FileHeader) {
	dynamicOSSConfig := global.ServerConfig.DynamicOSS
	url, err := servants.UploadFile(0, dynamicOSSConfig.Bucket, file.Filename, file, servants.FileTypeVideo)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, util.FulfillVideoOSSPrefix(util.TrimDomainPrefix(url)))
}
