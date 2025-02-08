package authrepo

import "database/sql"

// UserRepository implements UserRepo
type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (user_id, email, password, first_name, last_name, image, role, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, user.UserID, user.Email, user.Password, user.FirstName, user.LastName, user.Image, user.Role, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT user_id, email, password, first_name, last_name, image, role, created_at, updated_at FROM users WHERE email = ?`
	row := r.DB.QueryRow(query, email)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Image, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
