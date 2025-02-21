package authrepo

import (
	"errors"

	"forum/repositories/postrepo"
	"forum/repositories/shared"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserService manages user operations
type UserService struct {
	user        UserRepo
	shared      *shared.SharedConfig
	postService *postrepo.PostService
}

// NewUserService creates a new instance of UserService
func NewUserService(user UserRepo, posts *postrepo.PostService) *UserService {
	return &UserService{
		user:        user,
		shared:      shared.NewSharedConfig(),
		postService: posts,
	}
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
	userID, err := u.shared.GenerateUUID()
	if err != nil {
		return errors.New("oops, something went wrong. try again")
	}
	user.UserID = userID

	return u.user.CreateUser(user)
}

func (u *UserService) Login(email, password string) (*User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password cannot be empty")
	}

	// Retrieve user from database
	user, err := u.user.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("did you forget your password")
	}

	return user, nil
}

func (u *UserService) GetUserByID(user *User) (*User, error) {
	if user.UserID == "" {
		return nil, errors.New("bad request")
	}

	return u.user.GetUserByID(user)
}

func (u *UserService) UpdateUser(user *User) (*User, error) {
	if user.UserID == "" {
		return nil, errors.New("bad request")
	}

	if user.Image == "" {
		return nil, errors.New("bad request")
	}

	image := user.Image

	updatedUser, err := u.GetUserByID(user)
	if err != nil {
		return nil, errors.New("this user does not exist")
	}

	updatedUser.Image = image
	updatedUser.UpdatedAt = time.Now()

	return u.user.UpdateUser(updatedUser)
}

// GetUserDashboard retrieves user data
func (u *UserService) GetUserDashboard(user string) (*UserData, error) {
	if user == "" {
		return nil, errors.New("you need to login to access this resource")
	}

	var userData UserData

	posts, err := u.postService.GetPostsByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Posts = posts

	comments, err := u.postService.GetCommentsByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Comments = comments

	replies, err := u.postService.GetRepliesByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Replies = replies

	likes, err := u.postService.GetLikesByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Likes = likes

	dislikes, err := u.postService.GetDislikesByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Dislikes = dislikes

	return &userData, nil
}
