package authrepo

// UserRepo defines database operations for users
type UserRepo interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}
