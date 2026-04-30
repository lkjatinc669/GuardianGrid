package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteDB *sql.DB
	once     sync.Once
)

// InitSQLite initializes the SQLite database (singleton)
func InitSQLite(dbPath string) *sql.DB {
	once.Do(func() {
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatalf("failed to open sqlite DB: %v", err)
		}

		// Basic connection tuning
		db.SetMaxOpenConns(1) // SQLite works best with single writer
		db.SetMaxIdleConns(1)

		if err := db.Ping(); err != nil {
			log.Fatalf("failed to ping sqlite DB: %v", err)
		}

		sqliteDB = db

		initTables(sqliteDB)
	})

	return sqliteDB
}

// GetSQLite returns existing DB instance
func GetSQLite() *sql.DB {
	if sqliteDB == nil {
		log.Fatal("SQLite not initialized. Call InitSQLite first.")
	}
	return sqliteDB
}

// CloseSQLite closes DB connection
func CloseSQLite() {
	if sqliteDB != nil {
		sqliteDB.Close()
	}
}

// ----------------------------
// TABLE INITIALIZATION
// ----------------------------

func initTables(db *sql.DB) {
	createUsersTable(db)
	createAgentsTable(db)
}

// USERS TABLE (TOTP AUTH)
func createUsersTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		totp_secret TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}
}

// AGENTS TABLE (IDENTITY)
func createAgentsTable(db *sql.DB) {
	query := `
	
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		hostname TEXT,
		token TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create agents table: %v", err)
	}
}
