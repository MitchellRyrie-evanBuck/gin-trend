package system

import "github.com/afl-lxw/gin-trend/service"

type ApiGroup struct {
	JwtApi
}

var (
	systemUserConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemUserConfigService
)
