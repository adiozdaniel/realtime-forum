package commentrepo

import (
	"time"

	"forum/forumapp"
	"forum/repositories/shared"
)

// CommentRequest represents comment request
type CommmentRequest struct {
	PostID string `json:"post_id"`
}

// CommentByIDRequest represents comment like requests
type CommentByIDRequest struct {
	CommentID string `json:"comment_id"`
}

// Comment represents a comment in the database
type Comment struct {
	CommentID       string    `json:"comment_id"`
	PostID          string    `json:"post_id"`
	UserID          string    `json:"user_id"`
	Author          string    `json:"user_name"`
	ParentCommentID string    `json:"parent_comment_id,omitempty"`
	Content         string    `json:"content"`
	Likes           int       `json:"likes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// CommentRepo represents the repository for comments
type CommentRepo struct {
	app     *forumapp.ForumApp
	res     *shared.JSONRes
	comment CommentInterface
	shared  *shared.SharedConfig
}

// NewComment creates a new Comment instance
func NewCommentRepo(app *forumapp.ForumApp) *CommentRepo {
	commentRepo := NewCommentRepository(app.Db.Query)
	commentService := NewCommentService(commentRepo)
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()

	return &CommentRepo{
		app:     app,
		res:     res,
		comment: commentService,
		shared:  shared,
	}
}
