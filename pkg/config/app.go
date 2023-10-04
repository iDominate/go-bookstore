package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//"root:@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/simplerest?parseTime=true")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
