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

	// Configure connection pool
	d.Query.SetMaxOpenConns(10)
	d.Query.SetMaxIdleConns(5)

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
			user_name TEXT,
			image TEXT,
			role TEXT,
			bio TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,

		"posts": `
		CREATE TABLE IF NOT EXISTS posts (
			post_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			post_author TEXT NOT NULL,
			author_img TEXT,
			post_title TEXT NOT NULL,
			post_content TEXT NOT NULL,
			post_image TEXT,
			post_video TEXT,
			post_category TEXT NOT NULL,
			post_hasComments BOOLEAN DEFAULT TRUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_posts_post_id ON posts(post_id);
		CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);`,

		"comments": `
		CREATE TABLE IF NOT EXISTS comments (
			comment_id TEXT PRIMARY KEY,
			post_id TEXT NOT NULL,
			post_title TEXT,
			post_author TEXT,
			post_author_img TEXT,
			user_id TEXT NOT NULL,
			user_name TEXT NOT NULL,
			author_img TEXT,
			parent_comment_id TEXT,
			comment TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (parent_comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE
		);CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
		CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);`,

		"likes": `
		CREATE TABLE IF NOT EXISTS likes (
			like_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			post_id TEXT,
			comment_id TEXT,
			reply_id TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_likes_user_id ON likes(user_id);`,

		"dislikes": `
		CREATE TABLE IF NOT EXISTS dislikes (
			like_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			post_id TEXT,
			comment_id TEXT,
			reply_id TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_dislikes_user_id ON dislikes(user_id);`,

		"replies": `
		CREATE TABLE IF NOT EXISTS replies (
			reply_id TEXT PRIMARY KEY,
			comment_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			user_name TEXT NOT NULL,
			author_img TEXT NOT NULL,
			parent_reply_id TEXT,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (parent_reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE
		);`,

		"activities": `
		CREATE TABLE IF NOT EXISTS activities (
			activity_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			activity_type TEXT NOT NULL,
			activity_data TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		);`,

		"notifications": `
		CREATE TABLE IF NOT EXISTS notifications (
			notification_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			sender_id TEXT,
			post_id TEXT,
			comment_id TEXT,
			reply_id TEXT,
			like_id TEXT,
			dislike_id TEXT,
			notification_type TEXT NOT NULL,
			message TEXT,
			is_read BOOLEAN DEFAULT FALSE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (sender_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE,
			FOREIGN KEY (like_id) REFERENCES likes(like_id) ON DELETE CASCADE
		);`,
	}

	// Start a transaction
	tx, err := tm.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered from panic: %v", r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Enable foreign key support within transaction
	if _, err := tx.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	// Execute table creation queries
	for name, query := range tables {
		if _, err = tx.Exec(query); err != nil {
			return fmt.Errorf("failed to create %s table: %v", name, err)
		}
		log.Printf("Table %s created or already exists\n", name)
	}

	// Verify tables exist
	for name := range tables {
		var tableExists bool
		err := tx.QueryRow("SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?)", name).Scan(&tableExists)
		if err != nil || !tableExists {
			return fmt.Errorf("table %s was not created successfully", name)
		}
	}

	return nil
}

// Close properly closes the database connection
func (d *DataConfig) Close() {
	if d.Query != nil {
		if err := d.Query.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
