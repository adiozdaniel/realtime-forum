package postrepo

import (
	"errors"
	"time"

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

	if post.PostID == "" {
		post.PostID, _ = p.shared.GenerateUUID()
	}

	post.AuthorImg = "/static/profiles/" + post.UserID
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
		return nil, p.DeleteLike(like, "likes")
	}

	hasDisliked, err := p.post.HasUserDisliked(like.PostID, like.UserID, "Post")
	if err != nil {
		return nil, err
	}

	if hasDisliked != "" {
		like.LikeID = hasDisliked
		err := p.DeleteLike(like, "dislikes")
		if err != nil {
			return nil, err
		}
	}

	like.LikeID, _ = p.shared.GenerateUUID()
	like.CreatedAt = time.Now()
	return p.post.AddLike(like)
}

func (p *PostService) PostDisLike(dislike *Like) (*Like, error) {
	if dislike.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	haslike, err := p.post.HasUserLiked(dislike.PostID, dislike.UserID, "Post")
	if err != nil {
		return nil, err
	}

	if haslike != "" {
		dislike.LikeID = haslike
		err = p.DeleteLike(dislike, "likes")
		if err != nil {
			return nil, errors.New("failed to delete like")
		}
	}

	hasDisliked, err := p.post.HasUserDisliked(dislike.PostID, dislike.UserID, "Post")
	if err != nil {
		return nil, err
	}

	if hasDisliked != "" {
		dislike.LikeID = hasDisliked
		return nil, p.DeleteLike(dislike, "dislikes")
	}

	dislike.LikeID, _ = p.shared.GenerateUUID()
	dislike.CreatedAt = time.Now()
	return p.post.PostDislike(dislike)
}

func (p *PostService) CommentAddLike(like *Like) (*Like, error) {
	if like.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	haslike, err := p.post.HasUserLiked(like.CommentID, like.UserID, "Comment")
	if err != nil {
		return nil, err
	}

	if haslike != "" {
		like.LikeID = haslike
		return nil, p.DeleteLike(like, "likes")
	}

	like.LikeID, _ = p.shared.GenerateUUID()
	like.CreatedAt = time.Now()
	return p.post.AddLike(like)
}

func (p *PostService) DeleteLike(dislike *Like, entityType string) error {
	if dislike.LikeID == "" {
		return errors.New("like ID cannot be empty")
	}

	return p.post.DisLike(dislike, entityType)
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

	comment.AuthorImg = "/static/profiles/" + comment.UserID
	comment.CommentID, _ = p.shared.GenerateUUID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	return p.post.CreateComment(comment)
}

// CreateReply creates a new reply
func (p *PostService) CreateCommentReply(reply *Reply) (*Reply, error) {
	if reply.UserID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	if reply.CommentID == "" {
		return nil, errors.New("comment ID cannot be empty")
	}

	if reply.Content == "" {
		return nil, errors.New("reply content cannot be empty")
	}

	reply.AuthorImg = "/static/profiles/" + reply.UserID
	reply.ReplyID, _ = p.shared.GenerateUUID()
	reply.CreatedAt = time.Now()
	reply.UpdatedAt = time.Now()
	return p.post.CreateReply(reply)
}

// GetPostsByUserID retrieves all posts created by a specific user
func (p *PostService) GetPostsByUserID(userID string) ([]*Post, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return p.post.GetPostsByUserID(userID)
}

// GetCommentsByUserID retrieves all comments created by a specific user
func (p *PostService) GetCommentsByUserID(userID string) ([]*Comment, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return p.post.GetCommentsByUserID(userID)
}

// GetRepliesByUserID retrieves all replies created by a specific user
func (p *PostService) GetRepliesByUserID(userID string) ([]*Reply, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return p.post.GetRepliesByUserID(userID)
}

// GetLikesByUserID retrieves all likes created by a specific user
func (p *PostService) GetLikesByUserID(userID string) ([]*Like, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return p.post.GetLikesByUserID(userID)
}

// GetDislikesByUserID retrieves all dislikes created by a specific user
func (p *PostService) GetDislikesByUserID(userID string) ([]*Like, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	return p.post.GetDislikesByUserID(userID)
}
