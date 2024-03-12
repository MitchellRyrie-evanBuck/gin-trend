package router

import (
	"github.com/afl-lxw/gin-trend/router/system"
	"github.com/afl-lxw/gin-trend/router/user"
)

type RouterGroup struct {
	System system.RouterGroup
	User   user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
