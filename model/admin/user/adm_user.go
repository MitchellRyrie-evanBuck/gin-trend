package user

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
)

type AppTemplate struct {
	global.TREND_MODEL
	UserName string `gorm:"not null"`
	Password string `gorm:"type:text"`
	// 可以根据需求添加其他字段，例如创建时间、更新时间等
	UserID ksuid.KSUID `gorm:"type:varchar(255);index;comment:用户UUID"` // 外键，关联到用户表的 UUID
}

func (AppTemplate) TableName() string {
	return "admin_user"
}
