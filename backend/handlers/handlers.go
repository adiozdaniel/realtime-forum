package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Register user (dummy handler)
func(h *Repo) RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err := h.app.Db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", req.Username, req.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Respond
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, `{"message": "User registered successfully"}`)
}

// Login handler (dummy implementation)
func(h *Repo) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Login successful"}`)
}

// Posts handler (dummy implementation)
func(h *Repo) PostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"posts": []}`)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
