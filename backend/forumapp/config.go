package forumapp

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type ForumApp struct {
	Tmpls  *TemplateCache
	Db     *sql.DB
	Store  sync.Map
	Errors error
}

func newForumApp() *ForumApp {
	return &ForumApp{Tmpls: newTemplateCache()}
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

	// Check if the database is accessible
	err = f.Db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Optional: Create tables if they don't exist
	_, err = f.Db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
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

func (f *ForumApp) GenerateToken(userID int) (http.Cookie, error) {
	token, err := f.generateSecureToken()
	if err != nil {
		return http.Cookie{}, err
	}

	// Store the token in a session cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,             // Prevent JavaScript access
		Secure:   f.IsProduction(), // Secure in production
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	}

	// Store the token in the map
	f.Store.Store(userID, cookie)
	return *cookie, nil
}

// generateSecureToken creates a cryptographically secure random token
func (f *ForumApp) generateSecureToken() (string, error) {
	b := make([]byte, 16) // 16 bytes = 128 bits, similar to a UUID
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Format it as a UUID-like string
	token := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])

	return token, nil
}

// isProduction returns true if the server is running in production mode
func (f *ForumApp) IsProduction() bool {
	return os.Getenv("ENV") == "production"
}
