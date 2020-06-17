package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 数据库表格初始化
func Init() error {
	var (
		err    error
		dbtype = "sqlite3"
		dburl  = "app.db"
	)
	DB, err = gorm.Open(dbtype, dburl)
	if err != nil {
		return err
	}
	for _, m := range dbTables {
		if err := m.Migrate(DB); err != nil {
			return err
		}
	}

	return nil
}
