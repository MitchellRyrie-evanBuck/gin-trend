package initialize

import (
	"github.com/afl-lxw/gin-trend/docs"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/middleware"
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/afl-lxw/gin-trend/router"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func Routers() *gin.Engine {
	// 设置为发布模式
	if global.TREND_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery(), middleware.NoCache, middleware.Options, gzip.Gzip(gzip.DefaultCompression), gin.CustomRecovery(func(c *gin.Context, err any) {
		_ = response.FailWithDetailed(nil, cast.ToString(err), c)
		c.Abort()
	}))
	Router.Use(middleware.LimitHandler(tollbooth.NewLimiter(10, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Second,
	})))
	if global.TREND_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}
	//InstallPlugin(Router) // 安装插件

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/static")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面
	Router.StaticFS(global.TREND_CONFIG.Local.StorePath, http.Dir(global.TREND_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.TREND_LOG.Info("use middleware cors")

	docs.SwaggerInfo.BasePath = global.TREND_CONFIG.System.RouterPrefix
	Router.GET(global.TREND_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.TREND_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	//----------------------------ADMIN---------------------------------------------
	adminSystemRouter := router.RouteGroupApp.Admin.System.BaseRouter
	adminUserRouter := router.RouteGroupApp.Admin.User

	//----------------------------APP---------------------------------------------
	appSystemRouter := router.RouteGroupApp.App.System.BaseSystemRouter
	appWeChatRouter := router.RouteGroupApp.App.Wechat.BaseWeChatRouter
	appAuthRouter := router.RouteGroupApp.App.Auth.BaseAuthRouter

	//----------------------------------------------------------------------------

	PublicGroup := Router.Group(global.TREND_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.TREND_CONFIG.System.RouterPrefix)
	//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
	{
		// 健康监测
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		adminSystemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		// -----------------------------------------
		appSystemRouter.InitBaseRouter(PublicGroup)
		appWeChatRouter.InitBaseRouter(PublicGroup)
		appAuthRouter.InitBaseRouter(PublicGroup)
	}

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		adminUserRouter.InitBaseRouter(PublicGroup)
		// -----------------------------------------
	}
	return Router
}
