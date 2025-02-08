package repositories

import (
	"database/sql"
	"errors"
)

// CommentService manages comment operations
type CommentService struct {
	comment *CommentRepository
	db      *sql.DB
}

// NewCommentService creates a new instance of CommentService
func NewCommentService(db *sql.DB) *CommentService {
	comment := &CommentRepository{DB: db}
	return &CommentService{comment, db}
}

func (c *CommentService) CreateComment(comment *Comment) error {
	if comment.Comment == "" {
		return errors.New("comment cannot be empty")
	}
	if comment.UserID == "" {
		return errors.New("user ID cannot be empty")
	}
	if comment.PostID == "" {
		return errors.New("post ID cannot be empty")
	}

	return c.comment.CreateComment(comment)
}
