package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Database *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", "u1035350_jaskipin_front:jaskipin123@(45.130.230.90)/u1035350_jaskipin_front?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	Database = db
}
