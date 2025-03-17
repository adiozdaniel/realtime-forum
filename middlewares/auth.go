package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
)

// SetUserIDInContext adds the user ID to the request context
func (a *AuthContext) SetUserIDInContext(ctx context.Context, userID string) context.Context {
	newCtx := context.WithValue(ctx, userIDKey, userID)
	return newCtx
}

// GetUserIDFromContext retrieves the user ID from the request context
func (a *AuthContext) GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

// AuthMiddleware protects routes from unauthorized access and redirects if not logged in
func (a *AuthContext) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if session token exists
		cookie, err := r.Cookie("session_token")
		if err != nil {
			sendUnauthorizedResponse(w, "Unauthorized: you need to be logged in to access this resource")
			return
		}

		// Verify session in global store
		userID, found := a.getUserIDByToken(cookie.Value)
		if !found || !a.ValidateSession(userID, cookie.Value) {
			sendUnauthorizedResponse(w, "Unauthorized: session expired or replaced by a new login")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// getUserIDByToken searches for a user ID based on a session token.
func (a *AuthContext) getUserIDByToken(token string) (string, bool) {
	var userID string
	found := false

	a.Sessions.Range(func(key, value interface{}) bool {
		if storedToken, ok := value.(string); ok && storedToken == token {
			userID, _ = key.(string)
			found = true
			return false // Stop iteration
		}
		return true
	})

	return userID, found
}

// Helper function to send unauthorized JSON response
func sendUnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	var res JSONRes
	res.Err = true
	res.Message = message
	res.Data = nil

	// Encode payload into JSON
	out, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Oops, server misbehaving, try later!", http.StatusInternalServerError)
		return
	}

	// Set HTTP status code
	w.WriteHeader(http.StatusUnauthorized)

	// Write JSON response
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, "Oops, server misbehaving, try later!", http.StatusInternalServerError)
		return
	}
}
