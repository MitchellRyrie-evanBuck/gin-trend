package user

import (
	"github.com/afl-lxw/gin-trend/service/admin"
)

type ApiGroup struct {
	BaseUserApi
}

var (
	configUserService = admin.ServiceGroupApp.UserServiceGroup.BaseUserConfigService
)
