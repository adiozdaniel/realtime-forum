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
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, bio, created_at, updated_at FROM users WHERE email = ?`
	row := r.DB.QueryRow(query, email)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Ensure UserRepository implements all methods of UserRepo
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE user_id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) GetUserByID(user *User) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, bio, created_at, updated_at FROM users WHERE user_id = ?`
	row := r.DB.QueryRow(query, user.UserID)

	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.Bio, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *User) (*User, error) {
	query := `UPDATE users SET email = ?, password = ?, user_name = ?, image = ?, role = ?, bio = ?, updated_at = ? WHERE user_id = ?`
	_, err := r.DB.Exec(query, user.Email, user.Password, user.UserName, user.Image, user.Role, user.Bio, user.UpdatedAt, user.UserID)
	return user, err
}

// UsernameExists checks if a given username already exists in the database.
func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_name = ? LIMIT 1)"

	err := r.DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists, nil
}
