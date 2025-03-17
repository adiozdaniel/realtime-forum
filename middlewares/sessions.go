package middlewares

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"
)

// GenerateToken generates a new session token, replaces any previous session, and stores it in a session cookie
func (a *AuthContext) GenerateToken(userID string, w http.ResponseWriter) {
	a.mu.Lock()
	defer a.mu.Unlock()

	sessionToken := generateSessionToken()
	a.Sessions.Store(userID, sessionToken)

	// Create the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	})
}

// ValidateSession checks if the session token in the request matches the stored session
func (a *AuthContext) ValidateSession(userID, sessionToken string) bool {
	storedToken, ok := a.Sessions.Load(userID)
	if !ok {
		return false
	}

	tokenStr, _ := storedToken.(string) // Ensure correct type
	if tokenStr != sessionToken {
		return false
	}

	return true
}

// Logout removes the user's session
func (a *AuthContext) Logout(userID string) {
	a.Sessions.Delete(userID)
}

// generateSessionToken creates a cryptographically secure random session token
func generateSessionToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
