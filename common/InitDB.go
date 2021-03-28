package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"shen.com/ginvue/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(localhost)/GinVue?charset=utf8mb4&parseTime=True&loc=Local")
	DB = db
	if err != nil {
		log.Println("链接数据库出现问题，请检察")
		return DB
	}
	DB.AutoMigrate(&model.User{})
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
