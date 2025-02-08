package postrepo

/*
	These interfaces define the methods for interacting with the database.
*/

// PostRepository defines database operations for posts
type PostRepo interface {
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
	ListPosts() ([]*Post, error)
}
