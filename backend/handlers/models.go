package handlers

import (
	"forum/forumapp"
	"forum/response"
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

type Repo struct {
	app *forumapp.ForumApp
	res *response.JSONRes
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	return &Repo{app, response.NewJSONRes()}
}
