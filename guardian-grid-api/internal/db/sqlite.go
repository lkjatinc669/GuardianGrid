// internal/db/sqlite.go
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLite *gorm.DB

func InitSQLite() {
	db, err := gorm.Open(sqlite.Open("static/guardian.db"), &gorm.Config{})
	if err != nil {
		panic(err) // 👈 THIS LINE FIXES DEBUGGING
	}

	SQLite = db
}
