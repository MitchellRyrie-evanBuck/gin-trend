package common

import "github.com/afl-lxw/gin-trend/global"

type SysAreas struct {
	global.TREND_MODEL
	Title    string `json:"title" grm:"title:用户登录名"`
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`            // 用户登录密码
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
}
