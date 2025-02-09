package forumapp

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DataConfig manages the database connection
type DataConfig struct {
	Query *sql.DB
}

// TableManager handles table creation
type TableManager struct {
	db *sql.DB
}

// NewDb initializes the database connection
func NewDb() *DataConfig {
	return &DataConfig{}
}

// InitDB initializes the database connection and tables
func (d *DataConfig) InitDB(dbPath string) error {
	var err error

	// Open the SQLite database
	d.Query, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the database is accessible
	if err = d.Query.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Enable foreign keys
	if _, err = d.Query.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	// Create tables using TableManager
	tm := NewTableManager(d.Query)
	if err = tm.CreateTables(); err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

// NewTableManager initializes a TableManager instance
func NewTableManager(db *sql.DB) *TableManager {
	return &TableManager{db: db}
}

// CreateTables creates the necessary tables
func (tm *TableManager) CreateTables() error {
	tables := map[string]string{
		"users": `
			CREATE TABLE IF NOT EXISTS users (
				user_id TEXT PRIMARY KEY,
				email TEXT UNIQUE NOT NULL,
				password TEXT NOT NULL,
				first_name TEXT,
				last_name TEXT,
				image TEXT,
				role TEXT,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);`,
		"posts": `
			CREATE TABLE IF NOT EXISTS posts (
				post_id TEXT PRIMARY KEY,
				user_id TEXT NOT NULL,
				post_author TEXT NOT NULL,
				post_title TEXT NOT NULL,
				post_content TEXT NOT NULL,
				post_image TEXT,
				post_video TEXT,
				post_category TEXT NOT NULL,
				post_likes INTEGER DEFAULT 0,
				post_dislikes INTEGER DEFAULT 0,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
			);`,
		"comments": `
			CREATE TABLE IF NOT EXISTS comments (
				comment_id TEXT PRIMARY KEY,
				post_id TEXT NOT NULL,
				user_id TEXT NOT NULL,
				parent_comment_id TEXT,
				comment TEXT NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
				FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
				FOREIGN KEY (parent_comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE
			);`,
	}

	for name, query := range tables {
		if _, err := tm.db.Exec(query); err != nil {
			return fmt.Errorf("failed to create %s table: %v", name, err)
		}
		log.Printf("Table %s created or already exists\n", name)
	}

	return nil
}
