package system

import (
	"github.com/afl-lxw/gin-trend/api"
	"github.com/gin-gonic/gin"
)

type BaseSystemRouter struct {
}

func (t *BaseSystemRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("app")
	baseSystemApi := api.InterfaceConfig.App.SystemApiGroup.BaseSystemAreasAPI
	{
		baseRouter.GET("areas", baseSystemApi.GetSystemAreas)
	}
	return baseRouter
}
