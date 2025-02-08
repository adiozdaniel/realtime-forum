package repositories

/*
	These interfaces define the methods for interacting with the database.
*/

// CommentRepo defines database operations for comments
type CommentRepo interface {
	CreateComment(comment *Comment) error
	GetCommentByID(id string) (*Comment, error)
	DeleteComment(id string) error
	ListCommentsByPost(postID string) ([]*Comment, error)
}
