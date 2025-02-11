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

func (p *PostService) CreatePost(post *Post) error {
	if post.PostTitle == "" {
		return errors.New("post title cannot be empty")
	}
	if post.PostContent == "" {
		return errors.New("post content cannot be empty")
	}

	post.PostID, _ = p.shared.GenerateUUID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	if post.UserID == "" {
		return errors.New("user ID cannot be empty")
	}

	return p.post.CreatePost(post)
}

func (p *PostService) ListPosts() ([]*Post, error) {
	posts, err := p.post.ListPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) AddLike(postID string, userID string) (*Post, error) {
	if postID == "" {
		return nil, errors.New("post ID cannot be empty")
	}
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	err := p.post.AddLike(postID)
	if err != nil {
		return nil, err
	}

	post, err := p.post.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}
