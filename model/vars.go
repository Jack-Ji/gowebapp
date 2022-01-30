package model

import "gorm.io/gorm"

// 数据库连接
var DB *gorm.DB

// 单次访问数据库最大返回数量
const MAX_QUERY = 1000

// 数据模型
var dbTables = []migratable{
	User{},
}
