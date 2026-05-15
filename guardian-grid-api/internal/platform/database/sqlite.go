package database

import (
	"database/sql"
	"log"
	"sync"

	"guardian-grid-api/internal/platform/security"
	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteDB *sql.DB
	once     sync.Once
)

func InitSQLite() *sql.DB {
	once.Do(func() {
		dbPath := security.GetDBPath()
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatalf("failed to open sqlite DB: %v", err)
		}

		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)

		if err := db.Ping(); err != nil {
			log.Fatalf("failed to ping sqlite DB: %v", err)
		}

		sqliteDB = db
		initTables(sqliteDB)
	})

	return sqliteDB
}

func GetSQLite() *sql.DB {
	if sqliteDB == nil {
		return InitSQLite()
	}
	return sqliteDB
}

func CloseSQLite() {
	if sqliteDB != nil {
		sqliteDB.Close()
	}
}

func initTables(db *sql.DB) {
	createUsersTable(db)
	createAgentsTable(db)
}

func createUsersTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		totp_secret TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}
}

func createAgentsTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		hostname TEXT,
		token TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create agents table: %v", err)
	}
}
