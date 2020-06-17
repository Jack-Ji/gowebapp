package model

import "github.com/jinzhu/gorm"

// 数据模型迁移接口
type migratable interface {
	TableName() string
	Migrate(*gorm.DB) error
	Drop(*gorm.DB) error
}
