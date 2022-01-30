package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库表格初始化
func Init(dsn string) error {
	var err error

	cfg := gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gowebapp_",
			SingularTable: true,
		},
	}
	DB, err = gorm.Open(mysql.Open(dsn), &cfg)
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
