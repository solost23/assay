package controllers

import "github.com/gin-gonic/gin"

type ComputerController struct{}

func ComputerRegister(router *gin.RouterGroup) {
	controller := &ComputerController{}

	// 上位机
	router.GET("status", controller.status)
}

func (*ComputerController) status(c *gin.Context) {
	computerService.Status(c)
}
