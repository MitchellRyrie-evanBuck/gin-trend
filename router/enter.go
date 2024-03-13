package router

import (
	adminRouter "github.com/afl-lxw/gin-trend/router/admin"
	appRouter "github.com/afl-lxw/gin-trend/router/app"
)

type RouteConfigGroup struct {
	Admin adminRouter.RouterGroup
	App   appRouter.RouterGroup
}

var RouteGroupApp = new(RouteConfigGroup)
