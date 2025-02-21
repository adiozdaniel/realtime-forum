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
	PostDislike(dislike *Like) (*Like, error)

	// Shared
	ListPosts() ([]*Post, error)
	AddLike(like *Like) (*Like, error)
	DisLike(like *Like, entityType string) error
	HasUserLiked(entityID, userID string, entityType string) (string, error)
	HasUserDisliked(entityID, userID string, entityType string) (string, error)

	// Comments
	CreateComment(post *Comment) (*Comment, error)

	// Replies
	CreateReply(post *Reply) (*Reply, error)
}
