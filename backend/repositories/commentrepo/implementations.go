package commentrepo

import (
	"database/sql"
)

// CommentRepository implements CommentRepo
type CommentRepository struct {
	DB *sql.DB
}

// NewCommentRepository creates a new instance of CommentRepository
func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) CreateComment(comment *Comment) error {
	query := `INSERT INTO comments (comment_id, post_id, user_id, parent_comment_id, comment, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.CommentID, comment.PostID, comment.UserID, comment.ParentCommentID, comment.Comment, comment.CreatedAt, comment.UpdatedAt)
	return err
}

// GetCommentByID retrieves a comment by its ID
func (r *CommentRepository) GetCommentByID(id string) (*Comment, error) {
	query := `SELECT comment_id, post_id, user_id, parent_comment_id, comment, created_at, updated_at FROM comments WHERE comment_id = ?`
	row := r.DB.QueryRow(query, id)

	var comment Comment
	err := row.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment removes a comment by its ID
func (r *CommentRepository) DeleteComment(id string) error {
	query := `DELETE FROM comments WHERE comment_id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

// ListCommentsByPost retrieves all comments for a given post
func (r *CommentRepository) ListCommentsByPost(postID string) ([]*Comment, error) {
	query := `SELECT comment_id, post_id, user_id, parent_comment_id, comment, created_at, updated_at FROM comments WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
