package commentrepo

/*
	These interfaces define the methods for interacting with the database.
*/

// CommentInterface defines database operations for comments
type CommentInterface interface {
	CreateComment(comment *Comment) error
	GetCommentByID(id string) (*Comment, error)
	DeleteComment(id string) error
	ListCommentsByPost(postID string) ([]*Comment, error)
}
