package initialize

import (
	"os"

	"github.com/afl-lxw/gin-trend/global"
	//"github.com/afl-lxw/gin-trend/model/example"
	//"github.com/afl-lxw/gin-trend/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.TREND_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.TREND_DB
	err := db.AutoMigrate(

	//system.SysApi{},
	//system.SysUser{},
	//system.SysBaseMenu{},
	//system.JwtBlacklist{},
	//system.SysAuthority{},
	//system.SysDictionary{},
	//system.SysOperationRecord{},
	//system.SysAutoCodeHistory{},
	//system.SysDictionaryDetail{},
	//system.SysBaseMenuParameter{},
	//system.SysBaseMenuBtn{},
	//system.SysAuthorityBtn{},
	//system.SysAutoCode{},
	//system.SysExportTemplate{},
	//
	//example.ExaFile{},
	//example.ExaCustomer{},
	//example.ExaFileChunk{},
	//example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		global.TREND_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.TREND_LOG.Info("register table success")
}
