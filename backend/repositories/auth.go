package repositories

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

	// Parse user
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	// Query the database
	err := h.user.Register(&req)
	if err != nil {
		h.res.SetError(w, err, http.StatusConflict)
		return
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(req.UserID)

	// Set the session cookie
	http.SetCookie(w, &token)

	// Respond with success and token
	h.res.Data = req
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

	var userID, username, hashedPassword string

	err := h.app.Db.Query.QueryRowContext(ctx, "SELECT user_id, username, password FROM users WHERE email = ?", req.Email).
		Scan(&userID, &username, &hashedPassword)

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

	// Check if user already has a valid session
	if cookie, err := r.Cookie("session_token"); err == nil {
		if storedCookie, exists := h.app.Sessions.Load(userID); exists {
			if token, ok := storedCookie.(*http.Cookie); ok && token.Value == cookie.Value {
				h.res.Err = false
				h.res.Message = "Login successful (existing session)"
				h.res.Data = map[string]interface{}{
					"token":    token.Value,
					"user_id":  userID,
					"username": username,
				}
				h.res.WriteJSON(w, *h.res, http.StatusOK)
				return
			}
		}
	}

	// Generate a token (e.g., JWT)
	token := h.app.GenerateToken(userID)

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

// LogoutHandler handles user logout.
func (h *Repo) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.res.SetError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Get the session token from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			h.res.SetError(w, errors.New("no session found"), http.StatusUnauthorized)
			return
		}
		h.res.SetError(w, err, http.StatusBadRequest)
		return
	}

	sessionToken := cookie.Value

	h.app.Sessions.Delete(sessionToken)

	// Clear the session cookie by setting an expired cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	})

	// Respond with success
	h.res.Err = false
	h.res.Message = "Logout successful"

	err = h.res.WriteJSON(w, *h.res, http.StatusOK)
	if err != nil {
		h.res.SetError(w, err, http.StatusInternalServerError)
		return
	}
}

// CheckAuth confirms if a user is logged in
func (h *Repo) CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Return a success response if this protected route is reached
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"signedIn": true})
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
