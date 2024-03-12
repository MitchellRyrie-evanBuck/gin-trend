package system

import (
	"github.com/afl-lxw/gin-trend/service/admin"
)

type ApiGroup struct {
	JwtApi
	CapApi
}

var (
	configSystemUserService = admin.ServiceGroupApp.SystemServiceGroup.BaseSystemUserConfigService
)
