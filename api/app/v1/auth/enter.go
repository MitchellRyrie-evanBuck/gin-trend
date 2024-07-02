package auth

import "github.com/afl-lxw/gin-trend/service"

type ApiGroup struct {
	BaseGoogleApi
}

var (
	authGoogleService = service.ExportServiceConfig.App.AuthServiceGroup.GoogleAuthService
)
