package model

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/segmentio/ksuid"
)

type UserCard struct {
	global.TREND_MODEL
	UUID     ksuid.KSUID `json:"uuid" gorm:"type:varchar(255);index;comment:用户UUID"`
	Title    string      `json:"title" gorm:"comment:标题"`
	Content  string      `json:"content" gorm:"comment:内容"`
	ImgList  []string    `json:"imgList" gorm:"comment:图片列表"`
	Address  string      `json:"address" gorm:"comment:地址"`
	IsDelete bool        `json:"isDelete" gorm:"comment:是否删除"`
	IsHot    bool        `json:"isHot" gorm:"comment:推荐热门"`
	Tags     []Tags
}
