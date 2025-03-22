package authrepo

import (
	"database/sql"
	"errors"
)

// UserRepository implements UserRepo
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepo creates a new instance of UserRepo
func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (user_id, email, password, user_name, image, role, bio, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, user.UserID, user.Email, user.Password, user.UserName, user.Image, user.Role, user.Bio, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return errors.New("failed to create user, please try again later")
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, bio, created_at, updated_at FROM users WHERE email = ?`
	row := r.DB.QueryRow(query, email)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no user found with the given email")
		}
		return nil, errors.New("failed to retrieve user, please try again later")
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(userName string) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, bio, created_at, updated_at FROM users WHERE user_name = ?`
	row := r.DB.QueryRow(query, userName)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no user found with the given user name")
		}
		return nil, errors.New("failed to retrieve user, please try again later")
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE user_id = ?`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return errors.New("failed to delete user, please try again later")
	}
	return nil
}

func (r *UserRepository) GetUserByID(user *User) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, bio, created_at, updated_at FROM users WHERE user_id = ?`
	row := r.DB.QueryRow(query, user.UserID)

	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user does not exist")
		}
		return nil, errors.New("failed to retrieve user, please try again later")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *User) (*User, error) {
	query := `UPDATE users SET email = ?, password = ?, user_name = ?, image = ?, role = ?, bio = ?, updated_at = ? WHERE user_id = ?`
	_, err := r.DB.Exec(query, user.Email, user.Password, user.UserName, user.Image, user.Role, user.Bio, user.UpdatedAt, user.UserID)
	if err != nil {
		return nil, errors.New("failed to update user, please try again later")
	}
	return user, nil
}

// UsernameExists checks if a given username already exists in the database.
func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_name = ? LIMIT 1)"

	err := r.DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, errors.New("failed to check username availability, please try again later")
	}

	return exists, nil
}
