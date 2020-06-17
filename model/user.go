package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int64  `json:"primary_key;auto_increment"`
	Name     string `gorm:"unique_index;not null"`
	Password *string
	Salt     *string
	Icon     *string
}

func (User) TableName() string {
	return "gowebapp_user"
}

func (u User) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&u).Error
}

func (u User) Drop(db *gorm.DB) error {
	return db.DropTable(&u).Error
}

func (u *User) GetByName(name string) error {
	return DB.Where("name = ?", name).Take(u).Error
}
