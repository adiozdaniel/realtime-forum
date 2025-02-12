package authrepo

import (
	"database/sql"
	"time"
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
	query := `INSERT INTO users (user_id, email, password, user_name, image, role, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, user.UserID, user.Email, user.Password, user.UserName, user.Image, user.Role, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, created_at, updated_at FROM users WHERE email = ?`
	row := r.DB.QueryRow(query, email)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.CreatedAt, &user.UpdatedAt)
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

func (r *UserRepository) GetUserByID(id string) (*User, error) {
	query := `SELECT user_id, email, password, user_name, image, role, created_at, updated_at FROM users WHERE user_id = ?`
	row := r.DB.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.UserName, &user.Image, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *User) error {
	query := `UPDATE users SET email = ?, password = ?, user_name = ?, image = ?, role = ?, updated_at = ? WHERE user_id = ?`
	_, err := r.DB.Exec(query, user.Email, user.Password, user.UserName, user.Image, user.Role, time.Now(), user.UserID)
	return err
}
