package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"forum/repositories/authrepo"
)

type contextKey string

const userIDKey contextKey = "userID"

type AuthContext struct {
	res      *JSONRes
	Sessions *authrepo.Sessions
}

// JSONRes represents a JSON response structure.
type JSONRes struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewAuthContext(ses *authrepo.Sessions) *AuthContext {
	return &AuthContext{res: &JSONRes{}, Sessions: ses}
}

// SetUserIDInContext adds the user ID to the request context
func (a *AuthContext) SetUserIDInContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
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
		var userID string
		found := false

		a.Sessions.Sess.Range(func(key, value interface{}) bool {
			if token, ok := value.(*http.Cookie); ok && token.Value == cookie.Value {
				userID, _ = key.(string)
				found = true
				return false // Stop iteration
			}
			return true
		})

		if !found {
			sendUnauthorizedResponse(w, "Unauthorized: you need to be logged in to access this resource")
			return
		}

		// Set user ID in request context for handlers
		ctx := a.SetUserIDInContext(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
