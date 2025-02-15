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

func (p *PostService) PostAddLike(like *Like) (*Like, error) {
	if like.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	haslike, err := p.post.HasUserLiked(like.PostID, like.UserID, "Post")
	if err != nil {
		return nil, err
	}

	if haslike != "" {
		like.LikeID = haslike
		return nil, p.DisLike(like)
	}

	like.LikeID, _ = p.shared.GenerateUUID()
	like.CreatedAt = time.Now()
	return p.post.AddLike(like)
}

func (p *PostService) DisLike(dislike *Like) error {
	if dislike.LikeID == "" {
		return errors.New("like ID cannot be empty")
	}

	return p.post.DisLike(dislike)
}

// CreateComment creates a new comment
func (p *PostService) CreatePostComment(comment *Comment) (*Comment, error) {
	if comment.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	if comment.PostID == "" {
		return nil, errors.New("post ID cannot be empty")
	}

	if comment.Content == "" {
		return nil, errors.New("comment content cannot be empty")
	}

	comment.CommentID, _ = p.shared.GenerateUUID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	return p.post.CreateComment(comment)
}
