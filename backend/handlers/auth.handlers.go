package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles user registration.
func (h *Repo) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Username == "" {
		h.res.SetError(w, errors.New("username is required"), http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		h.res.SetError(w, errors.New("email is required"), http.StatusBadRequest)
		return
	}
	if req.Password == "" {
		h.res.SetError(w, errors.New("password is required"), http.StatusBadRequest)
		return
	}

	// Check if username or email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := h.app.Db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", req.Username, req.Email).Scan(&exists)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	if exists > 0 {
		h.res.SetError(w, errors.New("email already exists"), http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	// Save to database
	_, err = h.app.Db.ExecContext(ctx, "INSERT INTO users (username, email, password) VALUES (?, ?, ?)", req.Username, req.Email, string(hashedPassword))
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	h.res.Err = false
	h.res.Message = "User registered successfully"

	// Respond with JSON
	err = h.res.WriteJSON(w, *h.res, http.StatusCreated)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

// LoginHandler handles user login.
func (h *Repo) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Parse request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, errors.New("invalid request body"), http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Email == "" {
		h.res.SetError(w, errors.New("email is required"), http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		h.res.SetError(w, errors.New("password is required"), http.StatusBadRequest)
		return
	}

	// Fetch user from the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var (
		userID         int
		username       string
		hashedPassword string
	)
	err := h.app.Db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE email = ?", req.Email).Scan(&userID, &username, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.res.SetError(w, errors.New("user does not exist"), http.StatusUnauthorized)
			return
		}

		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		h.res.SetError(w, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	// Generate a token (e.g., JWT)
	token, err := h.app.GenerateToken(userID)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Err = false
	h.res.Message = "Login successful"
	h.res.Data = map[string]interface{}{
		"token":    token.Value,
		"user_id":  userID,
		"username": username,
	}

	err = h.res.WriteJSON(w, *h.res, http.StatusOK)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

// Posts handler (dummy implementation)
func (h *Repo) PostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"posts": []}`)
		return
	}

	h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
}
