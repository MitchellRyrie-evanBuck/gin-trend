package system

import (
	"github.com/afl-lxw/gin-trend/global"
)

type JwtBlacklist struct {
	global.TREND_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
