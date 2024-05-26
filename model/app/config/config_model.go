package config

import "github.com/afl-lxw/gin-trend/global"

type AppConfig struct {
	global.TREND_MODEL
	UserID uint // 这是外键
}
