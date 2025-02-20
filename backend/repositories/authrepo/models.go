package authrepo

import (
	"forum/forumapp"
	"forum/middlewares"
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
	app    *forumapp.ForumApp
	res    *shared.JSONRes
	user   *UserService
	shared *shared.SharedConfig
	auth   *middlewares.AuthContext
}

// NewAuthRepo creates a new instance of AuthRepo
func NewAuthRepo(app *forumapp.ForumApp, auth *middlewares.AuthContext) *AuthRepo {
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()
	userRepo := NewUserRepo(app.Db.Query)
	userService := NewUserService(userRepo)

	return &AuthRepo{
		app:    app,
		res:    res,
		auth:   auth,
		shared: shared,
		user:   userService,
	}
}
