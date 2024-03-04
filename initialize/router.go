package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(gin.Logger())

	//InstallPlugin(Router) // 安装插件
	//systemRouter := router.RouterGroupApp.System

	//PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	//PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
	return Router
}
