package v1

import (
	"github.com/afl-lxw/gin-trend/api/admin/v1/system"
	"github.com/afl-lxw/gin-trend/api/admin/v1/user"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupAdmin = new(ApiGroup)
