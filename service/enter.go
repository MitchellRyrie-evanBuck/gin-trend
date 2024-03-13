package service

import (
	"github.com/afl-lxw/gin-trend/service/admin"
	"github.com/afl-lxw/gin-trend/service/app"
)

type EnterConfig struct {
	Admin admin.ServiceGroup
	App   app.ServiceGroup
}

var ExportServiceConfig = new(EnterConfig)
