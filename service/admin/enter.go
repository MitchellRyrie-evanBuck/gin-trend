package admin

import (
	"github.com/afl-lxw/gin-trend/service/admin/system"
	"github.com/afl-lxw/gin-trend/service/admin/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
