package system

import "github.com/afl-lxw/gin-trend/service"

type ApiGroup struct {
	JwtApi
	CapApi
}

var (
	configSystemUserService = service.ServiceGroupApp.SystemServiceGroup.BaseSystemUserConfigService
)
