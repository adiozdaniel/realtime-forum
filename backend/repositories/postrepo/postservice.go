package postrepo

import (
	"errors"
	"forum/repositories/shared"
	"time"
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

func (p *PostService) CreatePost(post *Post) (*Post, error) {
	if post.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	if post.PostTitle == "" {
		return nil, errors.New("post title cannot be empty")
	}
	
	if post.PostContent == "" {
		return nil, errors.New("post content cannot be empty")
	}

	post.PostID, _ = p.shared.GenerateUUID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	post.HasComments = true

	return p.post.CreatePost(post)
}

func (p *PostService) ListPosts() ([]*Post, error) {
	posts, err := p.post.ListPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) AddLike(req *PostLike) (*PostLike, error) {
	if req.PostID == "" {
		return nil, errors.New("post ID cannot be empty")
	}
	if req.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	return p.post.AddLike(req)
}

func (p *PostService) DisLike(req *PostLike) (*PostLike, error) {
	if req.PostID == "" {
		return nil, errors.New("post ID cannot be empty")
	}
	if req.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	return p.post.DisLike(req)
}
