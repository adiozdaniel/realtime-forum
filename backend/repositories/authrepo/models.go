package authrepo

import (
	"database/sql"
	"forum/forumapp"
	"forum/response"
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
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Image     string    `json:"image,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthRepo represents the repository for authentication
type AuthRepo struct {
	app      *forumapp.ForumApp
	DB       *sql.DB
	res      *response.JSONRes
	user     *UserService
	Sessions *Sessions
}

// NewAuthRepo creates a new instance of AuthRepo
func NewAuthRepo(
	app *forumapp.ForumApp,
	db *sql.DB,
) *AuthRepo {
	res := response.NewJSONRes()
	user := NewUserService(db)
	return &AuthRepo{app, db, res, user, &Sessions{}}
}
