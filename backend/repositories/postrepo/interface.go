package postrepo

/*
	These interfaces define the methods for interacting with the database.
*/

// PostRepository defines database operations for posts
type PostRepo interface {
	CreatePost(post *Post) (*Post, error)
	GetPostByID(id string) (*Post, error)
	UpdatePost(post *Post) (*Post, error)
	DeletePost(id string) error
	ListPosts() ([]*Post, error)
	AddLike(req *PostLike) (*PostLike, error)
	DisLike(req *PostLike) (*PostLike, error)
}
