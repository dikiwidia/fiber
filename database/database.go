package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	var err error
	const MYSQL = "root:root@tcp(127.0.0.1:3306)/go_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := MYSQL
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database")
	}
	fmt.Println("Connected to Database")
}
