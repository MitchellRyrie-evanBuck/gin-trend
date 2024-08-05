package user

import (
	system "github.com/afl-lxw/gin-trend/model/admin/system"
)

type ServiceGroup struct {
	BaseUserConfigService
}

var (
	UserModel system.AdminSysUser
)
