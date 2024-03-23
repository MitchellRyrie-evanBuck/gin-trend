package user

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
)

type AppUser struct {
	global.TREND_MODEL
	UUID          ksuid.KSUID `json:"uuid" gorm:"type:varchar(255);index;comment:用户UUID"` // 用户UUID
	Username      string      `json:"userName" gorm:"index;comment:用户登录名"`                // 用户登录名
	Gender        int         `json:"gender" gorm:"default:0;comment: 性别"`
	Password      string      `json:"-"  gorm:"comment:用户登录密码"`                    // 用户登录密码
	NickName      string      `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`   // 用户昵称
	SideMode      string      `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"` // 用户侧边主题
	BackgroundImg string      `json:"backgroundImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	HeaderImg     string      `json:"headerImg" gorm:"default:dark;comment:用户头像背景"`
	BaseColor     string      `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`      // 基础颜色
	ActiveColor   string      `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"` // 活跃颜色
	Phone         string      `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email         string      `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable        int         `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Encryption    string      `json:"encryption" gorm:"comment: 密码加密"`
	BusinessCard
}

type BusinessCard struct {
	DouYin string `json:"douyin" gorm:"comment: 抖音名片" `
	QQ     string `json:"qq" gorm:"comment: qq名片" `
	WeChat string `json:"wechat" gorm:"comment: 微信名片" `
	Xhs    string `json:"xhs" gorm:"comment: 小红书名片" `
	WeBo   string `json:"webo" gorm:"comment: 微博名片" `
}

func (AppUser) TableName() string {
	return "app_users"
}
