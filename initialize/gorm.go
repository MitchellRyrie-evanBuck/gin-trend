package initialize

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/model/admin/token"
	"github.com/afl-lxw/gin-trend/model/app/message"
	"github.com/afl-lxw/gin-trend/model/app/template"
	"github.com/afl-lxw/gin-trend/utils"
	"log"
	"os"

	"github.com/afl-lxw/gin-trend/global"
	//"github.com/afl-lxw/gin-trend/model/example"
	"github.com/afl-lxw/gin-trend/model/admin/system"
	"github.com/afl-lxw/gin-trend/model/app/note"
	"github.com/afl-lxw/gin-trend/model/app/user"

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
		system.AdminSysUser{},
		user.AppUser{},
		note.AppNote{},
		template.AppTemplate{},
		token.AppToken{},
		message.AppMessage{},
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

func CheckTables() {
	db := global.TREND_DB
	// 检查表是否存在
	tableName := "areas"
	exists := utils.CheckTableExists(db, tableName)

	if exists {
		fmt.Printf("Table '%s' exists.\n", tableName)
	} else {
		fmt.Printf("Table '%s' does not exist. Importing SQL file...\n", tableName)
		// 导入 SQL 文件
		if err := utils.ImportSQLFile(db, "areas.sql"); err != nil {
			log.Fatal("Failed to import SQL file:", err)
		}
		fmt.Println("SQL file imported successfully.")
	}
}
