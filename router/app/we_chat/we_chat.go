package we_chat

import (
	"github.com/afl-lxw/gin-trend/api"
	"github.com/gin-gonic/gin"
)

type BaseWeChatRouter struct {
}

func (t *BaseWeChatRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("app")
	baseSystemApi := api.InterfaceConfig.App.WeChat.WeChart
	{
		baseRouter.GET("wechat", baseSystemApi.WXCheckSignature)
	}
	return baseRouter
}
