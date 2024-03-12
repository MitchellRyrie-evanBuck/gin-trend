package v1

import (
	"github.com/afl-lxw/gin-trend/api/v1/system"
	"github.com/afl-lxw/gin-trend/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
