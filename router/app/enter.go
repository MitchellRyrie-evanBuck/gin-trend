package app

import (
	"github.com/afl-lxw/gin-trend/router/app/system"
	"github.com/afl-lxw/gin-trend/router/app/we_chat"
)

type RouterGroup struct {
	System system.RouterGroup
	Wechat we_chat.RouterGroup
}
