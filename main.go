package main

import (
	"github.com/afl-lxw/gin-trend/core"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	//global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	//initialize.DBList()
	core.RunServer()
}
