// internal/db/mysql.go
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL *gorm.DB

func InitMySQL() {
	dsn := "user:password@tcp(127.0.0.1:3306)/guardian?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed mysql")
	}

	MySQL = db
}
