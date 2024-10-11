package services

import (
	"assay/infra/constant"
	"assay/infra/response"
	"assay/services/servants"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type OSSService struct{}

func (*OSSService) StaticUpload(c *gin.Context, file *multipart.FileHeader) {
	folder := "assay.static"

	url, err := servants.UploadFile(0, folder, file.Filename, file, servants.FileTypeImage)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, url)
}

func (*OSSService) DynamicUpload(c *gin.Context, file *multipart.FileHeader) {
	folder := "assay.dynamic"

	url, err := servants.UploadFile(0, folder, file.Filename, file, servants.FileTypeVideo)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	response.Success(c, url)
}
