package authrepo

import (
	"errors"
	"fmt"

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

	name := u.shared.CleanUsername(user.UserName)
	if name == "" {
		return errors.New("username cannot be empty and contains only letters")
	}

	user.UserName = name

	existingUser, _ := u.user.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("this email is already in use")
	}

	userName, _ := u.user.UsernameExists(user.UserName)
	if userName {
		return fmt.Errorf("%s username is already in use", user.UserName)
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

	go u.postService.RecordActivity(user.UserID, "account_registration", "registered an account")

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

	image := user.Image

	updatedUser, err := u.GetUserByID(user)
	if err != nil {
		return nil, errors.New("this user does not exist")
	}

	updatedUser.Image = image
	updatedUser.UpdatedAt = time.Now()

	var msg string

	if image != "" {
		msg = "updated user profile picture"
	} else {
		msg = "updated user profile"
	}

	go u.postService.RecordActivity(user.UserID, "profile", msg)

	return u.user.UpdateUser(updatedUser)
}

func (u *UserService) EditBio(user *User) (*User, error) {
	if user.UserID == "" {
		return nil, errors.New("bad request")
	}

	if user.Bio == "" {
		return nil, errors.New("bad request")
	}

	bio := user.Bio

	updatedUser, err := u.GetUserByID(user)
	if err != nil {
		return nil, errors.New("this user does not exist")
	}

	msg := "created a user bio"
	if updatedUser.Bio != bio {
		msg = "updated user bio"
	}

	updatedUser.Bio = bio
	updatedUser.UpdatedAt = time.Now()

	go u.postService.RecordActivity(user.UserID, "user_bio", msg)

	return u.user.UpdateUser(updatedUser)
}

// GetUserDashboard retrieves user data
func (u *UserService) GetUserDashboard(user string) (*UserData, error) {
	if user == "" {
		return nil, errors.New("you need to login to access this resource")
	}

	var userinfo User

	userinfo.UserID = user

	newUserInfo, err := u.user.GetUserByID(&userinfo)
	if err != nil {
		return nil, errors.New("oops something went wrong")
	}

	var userData UserData

	userData.UserInfo = newUserInfo
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

	activities, err := u.postService.GetActivitiesByUserID(user)
	if err != nil {
		return nil, err
	}
	userData.Activities = activities

	likedPosts, err := u.postService.GetLikedPosts(user)
	if err != nil {
		return nil, err
	}
	userData.LikedPosts = likedPosts

	return &userData, nil
}
