package routers

import (
	"assay/controllers"
	"github.com/gin-gonic/gin"
)

func InitImportRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("import")
	{
		controllers.ImportRegister(routerGroup)
	}
}
