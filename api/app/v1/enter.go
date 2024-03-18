package v1

import (
	"github.com/afl-lxw/gin-trend/api/app/v1/home_views"
	"github.com/afl-lxw/gin-trend/api/app/v1/search"
	"github.com/afl-lxw/gin-trend/api/app/v1/system"
	"github.com/afl-lxw/gin-trend/api/app/v1/user"
	"github.com/afl-lxw/gin-trend/api/app/v1/weChat"
)

type AppV1ApiGroup struct {
	SystemApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
	HomeViewGroup  home_views.ApiGroup
	SearchGroup    search.ApiGroup
	WeChat         weChat.ApiGroup
}

var ApiGroupAdmin = new(AppV1ApiGroup)
