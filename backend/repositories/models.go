package repositories

import (
	"forum/forumapp"
	"forum/response"
	"time"
)

type Repo struct {
	app     *forumapp.ForumApp
	res     *response.JSONRes
	comment *CommentService
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	// get the initialised db connection
	db := app.Db.Query

	// initialise services
	comment := NewCommentService(db)
	return &Repo{
		app, response.NewJSONRes(),
		comment,
	}
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
