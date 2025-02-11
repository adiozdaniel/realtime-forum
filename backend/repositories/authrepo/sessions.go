package authrepo

import (
	"net/http"
	"os"
	"sync"
	"time"
)

type Sessions struct {
	Sess sync.Map
}

// GenerateToken generates a new session token and stores it in a session cookie
func (s *Sessions) GenerateToken(userID string) http.Cookie {
	env := os.Getenv("ENV")

	// Store the token in a session cookie
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    userID,
		Path:     "/",
		HttpOnly: true,                // Prevent JavaScript access
		Secure:   env == "production", // Secure in production
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	}

	// Store the token in the map
	s.Sess.Store(userID, cookie)
	return *cookie
}
