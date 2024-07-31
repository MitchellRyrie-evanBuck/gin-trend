package config

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
)

type AppConfig struct {
	global.TREND_MODEL
	UserID ksuid.KSUID `gorm:"type:varchar(255);index;comment:用户UUID"` // 外键，关联到用户表的 UUID

}
