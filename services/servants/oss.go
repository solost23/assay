package servants

import (
	"assay/dao"
	"assay/infra/global"
	"assay/infra/util"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/minio/minio-go"
	"io"
	"mime/multipart"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	FileTypeImage = "image"
	FileTypeVideo = "video"
)

func UploadFile(userId uint, folderName string, fileName string, file *multipart.FileHeader, fileType string) (string, error) {
	fileHandle, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	b, err := io.ReadAll(fileHandle)
	if err != nil {
		return "", err
	}

	return Upload(userId, folderName, fileName, b, fileType)
}

func Upload(userId uint, folderName string, fileName string, fileBytes []byte, fileType string) (string, error) {
	if len(fileBytes) == 0 {
		return "", errors.New("upload file is empty")
	}
	mime := strings.Split(mimetype.Detect(fileBytes).String(), " ")[0]
	if !strings.HasPrefix(mime, fileType) {
		return "", errors.New(fmt.Sprintf("invalid mime type: %s", mime))
	}

	fileName = util.NewMd5(
		time.Now().Format(time.DateOnly)+
			fmt.Sprintf("%d", userId)+
			util.NewMd5(string(fileBytes))+
			fileName) + path.Ext(fileName)

	return upload(userId, folderName, fileName, fileBytes, fileType)
}

func upload(userId uint, folderName string, fileName string, fileBytes []byte, fileType string) (string, error) {
	if err := createBucket(folderName); err != nil {
		return "", err
	}

	reader := bytes.NewReader(fileBytes)
	client := global.Minio
	ctx := context.TODO()
	_, err := client.PutObjectWithContext(ctx, folderName, fileName, reader, reader.Size(), minio.PutObjectOptions{ContentType: fmt.Sprintf("Application/%s", fileType)})
	if err != nil {
		return "", err
	}

	// 生成url
	request := make(url.Values)
	fileUrl, err := client.PresignedGetObject(folderName, fileName, 168*time.Hour, request)
	if err != nil {
		return "", err
	}

	db := global.DB
	sqlOSSFile := &dao.OSSFile{
		Name:      fileName,
		Path:      folderName,
		Type:      fileType,
		Size:      reader.Size(),
		CreatorId: userId,
	}
	if err = dao.GInsert(db, sqlOSSFile); err != nil {
		return "", err
	}
	return fileUrl.String(), nil
}

func createBucket(bucketName string) error {
	client := global.Minio

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
	return nil
}
