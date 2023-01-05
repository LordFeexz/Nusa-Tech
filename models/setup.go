package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:qwertyui@/Nusa-Tech?charset=utf8&parseTime=true"))

	if err != nil {
		panic(err)
	} else {
		fmt.Println("success connect")
	}

	database.AutoMigrate(&User{})

	Db = database
}
