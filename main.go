package main

import (
	"github.com/afl-lxw/gin-trend/core"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/initialize"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 统一设置时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = location

	global.TREND_VP = core.Viper() // 初始化Viper
	global.TREND_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.TREND_LOG)
	global.TREND_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	//initialize.DBList()
	core.RunServer()
}
