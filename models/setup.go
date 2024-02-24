package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconn() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-rest-api-1"))
	if err != nil {
		panic("Couldn't open database connection")
	}

	db.AutoMigrate(&Product{})

	DB = db
}
