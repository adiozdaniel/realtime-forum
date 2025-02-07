package repositories

import (
	"forum/forumapp"
	"forum/response"
	"time"
)

// RegisterRequest represents the request body for user registration.
type Request struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse represents the response for user registration.
type Response struct {
	Message string `json:"message"`
	Error   bool   `json:"false"`
}

// LoginRequest represents the request body for user login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the response for user login.
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// LogoutRequest represents the request body for logout
type LogoutRequest struct {
	UserId string `json:"user_id"`
}

type Repo struct {
	app *forumapp.ForumApp
	res *response.JSONRes
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	return &Repo{app, response.NewJSONRes()}
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

// Post represents a post in the database
type Post struct {
	PostID       string    `json:"post_id"`
	UserID       string    `json:"user_id"`
	PostTitle    string    `json:"post_title"`
	PostContent  string    `json:"post_content"`
	PostImage    string    `json:"post_image,omitempty"`
	PostVideo    string    `json:"post_video,omitempty"`
	PostCategory string    `json:"post_category"`
	PostLikes    int       `json:"post_likes"`
	PostDislikes int       `json:"post_dislikes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Comment represents a comment in the database
type Comment struct {
	CommentID       string    `json:"comment_id"`
	PostID          string    `json:"post_id"`
	UserID          string    `json:"user_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty"`
	Comment         string    `json:"comment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
