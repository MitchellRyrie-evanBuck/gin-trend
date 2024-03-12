package admin

import (
	"github.com/afl-lxw/gin-trend/router/admin/system"
	"github.com/afl-lxw/gin-trend/router/admin/user"
)

type RouterGroup struct {
	System system.RouterGroup
	User   user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
