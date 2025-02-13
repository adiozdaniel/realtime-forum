package commentrepo

import (
	"errors"
	"forum/repositories/shared"
)

// CommentService manages comment operations
type CommentService struct {
	com    CommentInterface
	shared *shared.SharedConfig
}

// NewCommentService creates a new instance of CommentService
func NewCommentService(comment CommentInterface) *CommentService {
	return &CommentService{comment, shared.NewSharedConfig()}
}

func (c *CommentService) CreateComment(comment *Comment) ( *Comment, error) {
	if comment.Content == "" {
		return nil, errors.New("comment cannot be empty")
	}
	if comment.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	if comment.PostID == "" {
		return nil, errors.New("post ID cannot be empty")
	}

	return c.com.CreateComment(comment)
}

func (s *CommentService) GetCommentByID(comment *CommentByIDRequest) (*Comment, error) {
	return s.com.GetCommentByID(comment)
}

func (s *CommentService) DeleteComment(comment *CommentByIDRequest) error {
	return s.com.DeleteComment(comment)
}

func (s *CommentService) ListCommentsByPost(comment *CommmentRequest) ([]*Comment, error) {
	if comment.PostID == "" {
		return nil, errors.New("comment_id cannot be empty")
	}
	
	return s.com.ListCommentsByPost(comment)
}

func (s *CommentService) AddLike(like *CommentByIDRequest) error {
	if like.CommentID == "" {
		return errors.New("comment_id cannot be empty")
	}

	return s.com.AddLike(like)
}

func (s *CommentService) DisLike(like *CommentByIDRequest) error {
	if like.CommentID == "" {
		return errors.New("comment_id cannot be empty")
	}
	
	return s.com.DisLike(like)
}
