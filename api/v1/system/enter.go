package system

import "github.com/afl-lxw/gin-trend/service"

type ApiGroup struct {
	JwtApi
	CapApi
}

var (
	systemUserConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemUserConfigService
)
