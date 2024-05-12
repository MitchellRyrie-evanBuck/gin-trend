package message

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/app/user"
)

type AppMessage struct {
	global.TREND_MODEL
	UserId user.AppUser
}
