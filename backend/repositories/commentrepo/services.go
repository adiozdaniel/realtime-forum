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

	return c.com.CreateComment(comment)
}

func (s *CommentService) GetCommentByID(id string) (*Comment, error) {
	return s.com.GetCommentByID(id)
}

func (s *CommentService) DeleteComment(id string) error {
	return s.com.DeleteComment(id)
}

func (s *CommentService) ListCommentsByPost(postID string) ([]*Comment, error) {
	return s.com.ListCommentsByPost(postID)
}
