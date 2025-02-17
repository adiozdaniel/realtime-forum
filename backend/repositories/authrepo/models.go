package authrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
	"time"
)

// LoginRequest represents the request body for user login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LogoutRequest represents the request body for logout
type LogoutRequest struct {
	UserId string `json:"user_id"`
}

// User represents a user in the database
type User struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserName  string    `json:"user_name"`
	Image     string    `json:"image"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthRepo represents the repository for authentication
type AuthRepo struct {
	app      *forumapp.ForumApp
	res      *shared.JSONRes
	user     *UserService
	shared   *shared.SharedConfig
	Sessions *Sessions
}

// NewAuthRepo creates a new instance of AuthRepo
func NewAuthRepo(app *forumapp.ForumApp) *AuthRepo {
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()
	sessions := &Sessions{}
	userRepo := NewUserRepo(app.Db.Query)
	userService := NewUserService(userRepo)

	return &AuthRepo{
		app:      app,
		res:      res,
		user:     userService,
		shared:   shared,
		Sessions: sessions,
	}
}
