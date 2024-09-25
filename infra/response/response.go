package response

import (
	"assay/infra/global"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	SuccessCode       = 0
	responseKey       = "response"
	streamContentType = "application/octet-stream"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Msg     string `json:"message"`
	Success bool   `json:"success"`
}

type LogResponse struct {
	Response
	Stack any `json:"stack"`
}

func Error(c *gin.Context, code int, err error) {
	responseError(c, code, err)
}

func Success(c *gin.Context, data any) {
	responseSuccess(c, data)
}

func responseError(c *gin.Context, code int, err error) {
	stack := ""
	if global.ServerConfig.Mode != gin.ReleaseMode {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: code, Success: false, Msg: err.Error(), Data: ""}
	c.JSON(http.StatusOK, resp)
	logResp := &LogResponse{
		Response: *resp,
		Stack:    stack,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set(responseKey, string(logResponse))
	c.AbortWithError(http.StatusOK, err)
}

func responseSuccess(c *gin.Context, data any) {
	resp := &Response{Code: SuccessCode, Success: true, Msg: "", Data: data}
	c.JSON(http.StatusOK, resp)
	logResp := &LogResponse{
		Response: *resp,
	}
	logResponse, err := json.Marshal(logResp)
	if err != nil {
		zap.S().Error("error marshal response", err)
		return
	}
	c.Set(responseKey, string(logResponse))
}
