package commentrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
	"forum/response"
	"time"
)

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

// CommentRepo represents the repository for comments
type CommentRepo struct {
	app     *forumapp.ForumApp
	res     *response.JSONRes
	comment CommentInterface
	shared  *shared.SharedConfig
}

// NewComment creates a new Comment instance
func NewCommentRepo(app *forumapp.ForumApp) *CommentRepo {
	commentRepo := NewCommentRepository(app.Db.Query)
	commentService := NewCommentService(commentRepo)
	shared := shared.NewSharedConfig()

	return &CommentRepo{
		app:     app,
		res:     response.NewJSONRes(),
		comment: commentService,
		shared:  shared,
	}
}
