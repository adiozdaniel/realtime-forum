package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize database connection to SQLite
func InitDB() {
	var err error
	// Open the SQLite database (use a file path or :memory: for in-memory DB)
	db, err = sql.Open("sqlite3", "./forum.db") // SQLite database file
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Optional: Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

// Register user (dummy handler)
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save to database (use ? as placeholders for SQLite)
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", req.Username, req.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Respond
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, `{"message": "User registered successfully"}`)
}

// Login handler (dummy implementation)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Login successful"}`)
}

// Posts handler (dummy implementation)
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"posts": []}`)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
