package router

import "github.com/afl-lxw/gin-trend/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
