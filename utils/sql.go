package utils

import (
	"gorm.io/gorm"
	"os"
	"strings"
)

// CheckTableExists 检查表是否存在
func CheckTableExists(db *gorm.DB, tableName string) bool {
	return db.Migrator().HasTable(tableName)
}

// ImportSQLFile 导入 SQL 文件
func ImportSQLFile(db *gorm.DB, filename string) error {
	content, err := os.ReadFile("public/db/" + filename)
	if err != nil {
		return err
	}
	sqlArr := strings.Split(string(content), ";")
	for _, sql := range sqlArr {
		if sql == "" {
			continue
		}
		db.Exec(sql)
	}
	//return db.Exec(string(content)).Error
	return err
}
