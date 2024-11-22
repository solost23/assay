package initialize

import (
	"assay/infra/global"
	"fmt"

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

	// 创建存储桶
	policy := `{  
		"Version": "2012-10-17",  
		"Statement": [  
			{  
				"Effect": "Allow",  
				"Action": [  
					"s3:PutObject",  
					"s3:GetObject",  
					"s3:DeleteObject"  
				],  
				"Principal": "*",  
				"Resource": [  
					"arn:aws:s3:::%s/*"  
				]  
			}  
		]  
	}`
	staticOSSConfig := global.ServerConfig.StaticOSS
	if err = createBucket(client, staticOSSConfig.Bucket, fmt.Sprintf(policy, staticOSSConfig.Bucket)); err != nil {
		zap.S().Panic("failed create static_oss bucket: ", err)
	}
	dynamicOSSConfig := global.ServerConfig.DynamicOSS
	if err = createBucket(client, dynamicOSSConfig.Bucket, fmt.Sprintf(policy, dynamicOSSConfig.Bucket)); err != nil {
		zap.S().Panic("failed create dynamic_oss bucket: ", err)
	}

	global.Minio = client
}

func createBucket(client *minio.Client, bucketName string, policy string) error {
	exist, err := client.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	if err = client.MakeBucket(bucketName, ""); err != nil {
		return err
	}

	return client.SetBucketPolicy(bucketName, policy)
}
