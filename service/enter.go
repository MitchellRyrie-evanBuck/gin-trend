package service

import (
	"github.com/afl-lxw/gin-trend/service/system"
	"github.com/afl-lxw/gin-trend/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
