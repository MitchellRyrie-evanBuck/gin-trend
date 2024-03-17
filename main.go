package main

import (
	"database/sql"
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
	initialize.DBList()
	if global.TREND_DB != nil {
		initialize.RegisterTables() // 初始化表
		initialize.CheckTables()
		// 程序结束前关闭数据库链接
		db, _ := global.TREND_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	}
	core.RunServer()
}
