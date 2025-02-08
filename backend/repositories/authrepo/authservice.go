package authrepo

import (
	"errors"
	"forum/repositories/shared"

	"golang.org/x/crypto/bcrypt"
)

// UserService manages user operations
type UserService struct {
	user UserRepo
}

// NewUserService creates a new instance of UserService
func NewUserService(user UserRepo) *UserService {
	return &UserService{user}
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
	userID, err := shared.GenerateUUID()
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
