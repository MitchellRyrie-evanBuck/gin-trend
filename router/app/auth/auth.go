package auth

import (
	"github.com/afl-lxw/gin-trend/api"
	"github.com/gin-gonic/gin"
)

type BaseAuthRouter struct {
}

func (b *BaseAuthRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("app")
	baseAuthApi := api.InterfaceConfig.App.AuthApiGroup.BaseGoogleApi
	{
		baseRouter.POST("google", baseAuthApi.VerificationGoogleCode)
	}
	return baseRouter
}
