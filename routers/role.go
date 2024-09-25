package routers

import (
	"assay/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	routerGroup := Router.Group("roles")
	{
		controllers.RoleRegister(routerGroup)
	}
}
