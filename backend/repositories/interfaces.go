package repositories

/*
	These interfaces define the methods for interacting with the database.
*/

// PostRepo defines database operations for posts
type PostRepo interface {
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
	ListPosts() ([]*Post, error)
}

// CommentRepo defines database operations for comments
type CommentRepo interface {
	CreateComment(comment *Comment) error
	GetCommentByID(id string) (*Comment, error)
	DeleteComment(id string) error
	ListCommentsByPost(postID string) ([]*Comment, error)
}
