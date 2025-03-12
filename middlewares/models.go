package middlewares

import (
	"sync"

	"forum/forumapp"
	"forum/repositories/renders"
)

type contextKey string

const userIDKey contextKey = "userID"

type AuthContext struct {
	res      *JSONRes
	Sessions sync.Map
	mu       sync.Mutex
	app      *forumapp.ForumApp
	render   *renders.RendersRepo
}

// JSONRes represents a JSON response structure.
type JSONRes struct {
	Err     bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewAuthContext(app *forumapp.ForumApp) *AuthContext {
	render := renders.NewRendersRepo(app)
	return &AuthContext{
		res: &JSONRes{},
		app: app,
		render: render,
	}
}
