package commentrepo

/*
	These interfaces define the methods for interacting with the database.
*/

// CommentInterface defines database operations for comments
type CommentInterface interface {
	CreateComment(comment *Comment) (*Comment, error)
	GetCommentByID(comment *CommentByIDRequest) (*Comment, error)
	DeleteComment(comment *CommentByIDRequest) error
	ListCommentsByPost(comment *CommmentRequest) ([]*Comment, error)
	AddLike(like *CommentByIDRequest) error
	DisLike(like *CommentByIDRequest) error
}
