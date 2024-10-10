package initialize

import (
	"assay/infra/global"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
)

func initMinio() {
	minioConfig := global.ServerConfig.Minio

	client, err := minio.New(
		minioConfig.EndPoint,
		minioConfig.SecretAccessKey,
		minioConfig.AccessKeyId,
		minioConfig.UserSsl,
	)
	if err != nil {
		zap.S().Panic("failed to connect minio: ", err)
	}

	global.Minio = client
}
