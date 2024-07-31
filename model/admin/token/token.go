package token

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
	"time"
)

type AppToken struct {
	global.TREND_MODEL
	Token     string `gorm:"not null"`
	ExpiredAt time.Time
	Type      int         `gorm:"type:integer"`
	UserID    ksuid.KSUID `gorm:"type:varchar(255);index;comment:用户UUID"` // 外键，关联到用户表的 UUID
}

func (AppToken) TableName() string {
	return "admin_token"
}
