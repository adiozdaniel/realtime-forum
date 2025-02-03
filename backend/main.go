package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3" 
)

var db *sql.DB

func main() {

	initDB() // Initialize SQLite database connection

	// Create server
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/api/auth/register", registerHandler)
	mux.HandleFunc("/api/auth/login", loginHandler)
	mux.HandleFunc("/api/posts", postsHandler)

	// CORS middleware
	handler := corsMiddleware(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting @http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}

// Initialize database connection to SQLite
func initDB() {
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
func registerHandler(w http.ResponseWriter, r *http.Request) {
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
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Login successful"}`)
}

// Posts handler (dummy implementation)
func postsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"posts": []}`)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
