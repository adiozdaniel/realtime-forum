package repositories

import (
	"errors"
)

// UserService manages user operations
type UserService struct {
	user *UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService() *UserService {
	user := &UserRepository{}
	return &UserService{user}
}

func (u *UserService) Register(user *User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password cannot be empty")
	}

	existingUser, _ := u.user.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("this email is already in use")
	}

	// Hash the password (assuming a HashPassword function exists)
	hashedPassword := user.Password // TODO: Replace with actual hashing function
	user.Password = hashedPassword

	return u.user.CreateUser(user)
}

// PostService manages post operations
type PostService struct {
	post *PostRepository
}

// NewPostService creates a new instance of PostService
func NewPostService() *PostService {
	post := &PostRepository{}
	return &PostService{post}
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
}

// NewCommentService creates a new instance of CommentService
func NewCommentService() *CommentService {
	comment := &CommentRepository{}
	return &CommentService{comment}
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
