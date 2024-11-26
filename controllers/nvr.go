package controllers

import "github.com/gin-gonic/gin"

type NvrController struct{}

func NvrRegister(router *gin.RouterGroup) {
	controller := &NvrController{}

	// 通道列表
	router.GET("channel", controller.channel)
}

func (*NvrController) channel(c *gin.Context) {
	nvrService.Channel(c)
}
