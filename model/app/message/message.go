package message

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/app/user"
)

type AppMessage struct {
	global.TREND_MODEL
	UserId         user.AppUser
	OppositeUserId string
	SenderID       uint   `gorm:"not null; comment:发送者ID" json:"sender_id"`
	ReceiverID     uint   `gorm:"not null; comment:接收者ID" json:"receiver_id"`
	Content        string `gorm:"type:text;not null; comment:消息的内容" json:"content"`
}
