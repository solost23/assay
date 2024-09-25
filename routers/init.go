package routers

import (
	"assay/middlewares"

	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	apiGroup := router.Group("api/assay")
	{
		InitLoginRouter(apiGroup)
	}
	apiGroup.Use(
		middlewares.JWTAuth(),
	)
	{
		InitRoleRouter(apiGroup)
		InitUserRouter(apiGroup)
	}

	return router
}
