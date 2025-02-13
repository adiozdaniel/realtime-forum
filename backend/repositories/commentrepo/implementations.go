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

func (r *CommentRepository) CreateComment(comment *Comment) (*Comment, error) {
	query := `INSERT INTO comments (comment_id, post_id, user_id, user_name, parent_comment_id, content, likes, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.CommentID, comment.PostID, comment.UserID, &comment.Author, comment.ParentCommentID, comment.Content, comment.Likes, comment.CreatedAt, comment.UpdatedAt)
	return comment, err
}

// GetCommentByID retrieves a comment by its ID
func (r *CommentRepository) GetCommentByID(comm *CommentByIDRequest) (*Comment, error) {
	query := `SELECT comment_id, post_id, user_id, user_name, parent_comment_id, content, likes, created_at, updated_at FROM comments WHERE comment_id = ?`
	row := r.DB.QueryRow(query, comm.CommentID)

	var comment Comment
	err := row.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Author, &comment.ParentCommentID, &comment.Content, &comment.Likes, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment removes a comment by its ID
func (r *CommentRepository) DeleteComment(comm *CommentByIDRequest) error {
	query := `DELETE FROM comments WHERE comment_id = ?`
	_, err := r.DB.Exec(query, comm.CommentID)
	return err
}

// ListCommentsByPost retrieves all comments for a given post
func (r *CommentRepository) ListCommentsByPost(comment *CommmentRequest) ([]*Comment, error) {
	query := `SELECT comment_id, post_id, user_id, user_name, parent_comment_id, content, likes, created_at, updated_at FROM comments WHERE post_id = ?`
	rows, err := r.DB.Query(query, comment.PostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Author, &comment.ParentCommentID, &comment.Content, &comment.Likes, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// AddLike increments the like count for a comment
func (r *CommentRepository) AddLike(like *CommentByIDRequest) error {
	query := `UPDATE comments SET likes = likes + 1 WHERE comment_id = ?`
	_, err := r.DB.Exec(query, like.CommentID)
	return err
}

// DisLike decrements the like count for a comment
func (r *CommentRepository) DisLike(like *CommentByIDRequest) error {
	query := `UPDATE comments SET likes = likes - 1 WHERE comment_id = ?`
	_, err := r.DB.Exec(query, like.CommentID)
	return err
}
