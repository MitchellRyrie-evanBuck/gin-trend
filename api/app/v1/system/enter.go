package system

import "github.com/afl-lxw/gin-trend/service"

type ApiGroup struct {
	BaseSystemAreasAPI
}

var (
	configSystemUserService = service.ExportServiceConfig.App.SystemServiceGroup.BaseSystemAreas
)
