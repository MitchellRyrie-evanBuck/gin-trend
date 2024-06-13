package user

import (
	"github.com/afl-lxw/gin-trend/global"
	app "github.com/afl-lxw/gin-trend/model/app/config"
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
	DescriptionInfo
	Config   app.AppConfig
	ConfigID uint // 这是外键
	//BlackList []BlackList
	//WhiteList []WhiteList
}

type BlackList struct {
	UserId string `json:"userId" gorm:"type:varchar(255);index;comment:用户ID"`
}

type WhiteList struct {
	UserId string `json:"userId" gorm:"type:varchar(255);index;comment:用户ID"`
}

type BusinessCard struct {
	DouYin string `json:"douyin" gorm:"comment: 抖音名片" `
	QQ     string `json:"qq" gorm:"comment: qq名片" `
	WeChat string `json:"wechat" gorm:"comment: 微信名片" `
	Xhs    string `json:"xhs" gorm:"comment: 小红书名片" `
	WeBo   string `json:"webo" gorm:"comment: 微博名片" `
}

type DescriptionInfo struct {
	Intro                string `json:"u_intro" gorm:"type:varchar(300);default:null;comment:简介"`
	HeadPortrait         string `json:"u_head_portrait" gorm:"type:varchar(100);comment:头像"`
	Age                  int    `json:"u_age" gorm:"default:null;comment:年龄"`
	Constellation        string `json:"u_constellation" gorm:"type:char(6);default:null;comment:星座"`
	BloodType            string `json:"u_blood_type" gorm:"type:varchar(10);default:null;comment:血型"`
	SchoolTag            string `json:"u_school_tag" gorm:"type:varchar(50);default:null;comment:毕业学校"`
	Vocation             string `json:"u_vocation" gorm:"type:varchar(30);default:null;comment:职业"`
	NationID             int    `json:"u_nation_id" gorm:"default:null;comment:国家ID;foreign_key"`
	ProvinceID           int    `json:"u_province_id" gorm:"default:null;comment:省份ID;foreign_key"`
	CityID               int    `json:"u_city_id" gorm:"default:null;comment:城市ID;foreign_key"`
	FriendshipPolicyID   int    `json:"u_friendship_policy_id" gorm:"default:null;comment:好友策略ID;foreign_key"`
	UserStateID          int    `json:"u_user_state_id" gorm:"default:null;comment:用户状态ID;foreign_key"`
	FriendPolicyQuestion string `json:"u_friend_policy_question" gorm:"type:varchar(30);default:null;comment:好友策略问题"`
	FriendPolicyAnswer   string `json:"u_friend_policy_answer" gorm:"type:varchar(30);default:null;comment:好友策略答案"`
}

func (AppUser) TableName() string {
	return "app_users"
}
