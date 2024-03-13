package api

import (
	adminV1 "github.com/afl-lxw/gin-trend/api/admin/v1"
	appV1 "github.com/afl-lxw/gin-trend/api/app/v1"
)

type InterfaceApiConfig struct {
	Admin adminV1.AdminV1ApiGroup
	App   appV1.AppV1ApiGroup
}

var InterfaceConfig = new(InterfaceApiConfig)
