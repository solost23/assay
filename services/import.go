package services

import (
	"assay/infra/constant"
	"assay/infra/response"
	"assay/services/servants"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"io"
	"mime/multipart"
	"strings"
)

type ImportService struct{}

func (*ImportService) TemplateTasks(c *gin.Context) {
	xlsx := excelize.NewFile()

	_ = xlsx.SetSheetRow(servants.SheetName, "A1", &servants.TaskTitles)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"任务导入模板.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func (*ImportService) UploadTasks(c *gin.Context, file *multipart.FileHeader) {
	fileHandle, err := file.Open()
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer fileHandle.Close()

	fileByte, err := io.ReadAll(fileHandle)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}

	xlsx, err := excelize.OpenReader(strings.NewReader(string(fileByte)))
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	defer xlsx.Close()

	if xlsx.SheetCount > 1 {
		response.Error(c, constant.BadRequestCode, errors.New("只允许有一个Sheet"))
		return
	}

	rows, err := xlsx.GetRows(xlsx.WorkBook.Sheets.Sheet[0].Name)
	if err != nil {
		response.Error(c, constant.InternalServerErrorCode, err)
		return
	}
	fmt.Println(rows)
	// 检查数据合法性
	// 存储数据
	response.Success(c, "success")
}
