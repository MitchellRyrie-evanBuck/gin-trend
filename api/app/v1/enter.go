package v1

import (
	"github.com/afl-lxw/gin-trend/api/app/v1/home_views"
	"github.com/afl-lxw/gin-trend/api/app/v1/search"
)

type AppV1ApiGroup struct {
	SystemApiGroup home_views.ApiGroup
	UserApiGroup   search.ApiGroup
}

var ApiGroupAdmin = new(AppV1ApiGroup)
