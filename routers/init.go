package routers

import "github.com/gin-gonic/gin"

func Register() *gin.Engine {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	apiGroup := router.Group("api/assay")
	apiGroup.Use()
	{
		InitRoleRouter(apiGroup)
		InitUserRouter(apiGroup)
	}

	return router
}
