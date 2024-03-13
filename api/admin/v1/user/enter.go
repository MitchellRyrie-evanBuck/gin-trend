package user

import (
	"github.com/afl-lxw/gin-trend/service"
)

type ApiGroup struct {
	BaseUserApi
}

var (
	configUserService = service.ExportServiceConfig.Admin.UserServiceGroup.BaseUserConfigService
)
