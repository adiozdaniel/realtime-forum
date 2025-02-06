package forumapp

import (
	"database/sql"
	"fmt"
	"sync"
)

type ForumApp struct {
	Tmpls  *TemplateCache
	Db     *sql.DB
	Errors error
}

func newForumApp() *ForumApp {
	return &ForumApp{
		Tmpls: newTemplateCache(),
	}
}

var (
	instance *ForumApp
	once     sync.Once
)

func ForumInit() (*ForumApp, error) {
	var err error
	once.Do(func() {
		instance = newForumApp()
	})

	err = instance.InitDB()
	if err != nil {
		return nil, err
	}

	err = instance.Tmpls.CreateTemplatesCache()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// Initialize database connection to SQLite
func (f *ForumApp) InitDB() error {
	var err error
	// Open the SQLite database (use a file path or :memory: for in-memory DB)
	f.Db, err = sql.Open("sqlite3", "./forum.db") // SQLite database file
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Optional: Create tables if they don't exist
	_, err = f.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return nil
}
