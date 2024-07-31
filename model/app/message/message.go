package message

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
)

type AppMessage struct {
	global.TREND_MODEL
	UserID ksuid.KSUID `gorm:"type:varchar(255);index;comment:用户UUID"` // 外键，关联到用户表的 UUID

	OppositeUserId string
	SenderID       uint   `gorm:"not null; comment:发送者ID" json:"sender_id"`
	ReceiverID     uint   `gorm:"not null; comment:接收者ID" json:"receiver_id"`
	Content        string `gorm:"type:text;not null; comment:消息的内容" json:"content"`
}

func (AppMessage) TableName() string {
	return "app_message"
}
