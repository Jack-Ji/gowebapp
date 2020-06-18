package model

import (
	"gowebapp/utils"
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int64  `json:"primary_key;auto_increment"`
	Name     string `gorm:"unique_index;not null"`
	Password *string
	Salt     *string
}

func (User) TableName() string {
	return "gowebapp_user"
}

func (u User) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&u).Error
	if err != nil {
		return err
	}

	// 初始化admin用户
	var cnt int
	err = DB.Model(u).Where("name = ?", "admin").Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt == 0 {
		salted, salt := utils.GenSaltedPasswd("admin123456")
		u.Name = "admin"
		u.Password = &salted
		u.Salt = &salt
		err = DB.Create(&u).Error
		if err != nil {
			return err
		}
		log.Printf("已初始化admin账户, 密码为admin123456 ID为%d", u.ID)
	}

	return nil
}

func (u User) Drop(db *gorm.DB) error {
	return db.DropTable(&u).Error
}

func (u *User) GetByName(name string) error {
	err := DB.Where("name = ?", name).Take(u).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		err = ErrNotExist
	}
	return err
}
