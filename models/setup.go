package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:qwertyui@tcp(localhost:3000)/Nusa-Tech?charset=utf8&parseTime=true"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
}
