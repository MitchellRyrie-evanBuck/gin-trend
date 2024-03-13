package user

import (
	"github.com/afl-lxw/gin-trend/api"

	"github.com/gin-gonic/gin"
)

type BaseUserRouter struct {
}

func (b *BaseUserRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("trend")
	baseUserApi := api.InterfaceConfig.Admin.UserApiGroup.BaseUserApi
	{
		baseRouter.POST("user", baseUserApi.UserCreatedAPIFn)
		baseRouter.PUT("user", baseUserApi.UserUpdateAPIFn)
	}
	return baseRouter
}
