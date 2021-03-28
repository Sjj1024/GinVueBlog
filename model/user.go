package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(20)"`
	Phone    string
}
