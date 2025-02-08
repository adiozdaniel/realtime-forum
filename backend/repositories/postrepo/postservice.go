package postrepo

import (
	"errors"
	"forum/repositories/shared"
)

// PostService manages post operations
type PostService struct {
	post   PostRepo
	shared *shared.SharedConfig
}

// NewPostService creates a new instance of PostService
func NewPostService(post PostRepo) *PostService {
	shared := shared.NewSharedConfig()
	return &PostService{post, shared}
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
