package repositories

import (
	"database/sql"
)

/*
	These concrete implementations of the interfaces interact with the database.
*/

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

// PostRepository implements PostRepo
type PostRepository struct {
	DB *sql.DB
}

func (r *PostRepository) CreatePost(post *Post) error {
	query := `INSERT INTO posts (post_id, user_id, post_title, post_content, post_image, post_video, post_category, post_likes, post_dislikes, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, post.PostID, post.UserID, post.PostTitle, post.PostContent, post.PostImage, post.PostVideo, post.PostCategory, post.PostLikes, post.PostDislikes, post.CreatedAt, post.UpdatedAt)
	return err
}

// CommentRepository implements CommentRepo
type CommentRepository struct {
	DB *sql.DB
}

func (r *CommentRepository) CreateComment(comment *Comment) error {
	query := `INSERT INTO comments (comment_id, post_id, user_id, parent_comment_id, comment, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.CommentID, comment.PostID, comment.UserID, comment.ParentCommentID, comment.Comment, comment.CreatedAt, comment.UpdatedAt)
	return err
}
