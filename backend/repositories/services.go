package repositories

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserService manages user operations
type UserService struct {
	user *UserRepository
	db   *sql.DB
}

// NewUserService creates a new instance of UserService
func NewUserService(db *sql.DB) *UserService {
	user := &UserRepository{DB: db}
	return &UserService{user, db}
}

func (u *UserService) Register(user *User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email or password cannot be empty")
	}

	existingUser, _ := u.user.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("this email is already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("oops, something went wrong. try again")
	}
	user.Password = string(hashedPassword)

	// Generate a unique user ID
	userID, err := generateUUID()
	if err != nil {
		return errors.New("oops, something went wrong. try again")
	}
	user.UserID = userID

	return u.user.CreateUser(user)
}

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

// generateUUID creates a cryptographically secure random token
func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Format it as a UUID-like string
	token := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])

	return token, nil
}
