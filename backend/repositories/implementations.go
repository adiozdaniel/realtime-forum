package repositories

import (
	"database/sql"
)

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
