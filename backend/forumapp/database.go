package forumapp

import (
	"database/sql"
	"fmt"
)

// DataConfig represents the database configuration
type DataConfig struct {
	Db *sql.DB
}

// NewDb initializes the database connection
func NewDb() *DataConfig {
	return &DataConfig{}
}

// Initialize database connection to SQLite
func (d *DataConfig) InitDB() error {
	var err error
	// Open the SQLite database (use a file path or :memory: for in-memory DB)
	d.Db, err = sql.Open("sqlite3", "./forum.db") // SQLite database file
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the database is accessible
	err = d.Db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Optional: Create tables if they don't exist
	_, err = d.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}
