package middlewares

import "sync"

type contextKey string

const userIDKey contextKey = "userID"

type AuthContext struct {
	res      *JSONRes
	Sessions sync.Map
	mu       sync.Mutex
}

// JSONRes represents a JSON response structure.
type JSONRes struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewAuthContext() *AuthContext {
	return &AuthContext{res: &JSONRes{}}
}
