package repositories

import (
	"database/sql"
	"errors"
)

// PostService manages post operations
type PostService struct {
	post *PostRepository
	db   *sql.DB
}

// NewPostService creates a new instance of PostService
func NewPostService(db *sql.DB) *PostService {
	post := &PostRepository{DB: db}
	return &PostService{post, db}
}

func (p *PostService) CreatePost(post *Post) error {
	if post.PostTitle == "" {
		return errors.New("post title cannot be empty")
	}
	if post.PostContent == "" {
		return errors.New("post content cannot be empty")
	}
	if post.UserID == "" {
		return errors.New("user ID cannot be empty")
	}

	return p.post.CreatePost(post)
}

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
