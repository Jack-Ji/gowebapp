package model

import "gorm.io/gorm"

// 数据模型迁移接口
type migratable interface {
	Migrate(*gorm.DB) error
	Drop(*gorm.DB) error
}
