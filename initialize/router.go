package initialize

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(gin.Logger())
	Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	//InstallPlugin(Router) // 安装插件
	//systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	//PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	systemRouter := router.RouterGroupApp.System

	//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	return Router
}
