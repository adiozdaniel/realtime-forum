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
	GetPostsByUserID(userID string) ([]*Post, error)
	GetCommentsByUserID(userID string) ([]*Comment, error)
	GetRepliesByUserID(userID string) ([]*Reply, error)
	GetLikesByUserID(userID string) ([]*Like, error)
	GetLikedPostsByUserID(userID string) ([]*Post, error)
	GetDislikesByUserID(userID string) ([]*Like, error)
	AddActivity(activity *Activity) (*Activity, error)
	GetActivitiesByUserID(userID string) ([]*Activity, error)
	GetPostByLikeID(likeID, userID string) (*Post, error)
	GetCommentByLikeID(likeID, userID string) (*Comment, error)

	// Comments
	CreateComment(post *Comment) (*Comment, error)
	GetCommentByID(id string) (*Comment, error)
	UpdateComment(comment *Comment) (*Comment, error)
	DeleteComment(id string) error

	// Replies
	CreateReply(post *Reply) (*Reply, error)
	GetReplyByID(id string) (*Reply, error)

	// Notifications
	CreateNotification(n *Notification) (*Notification, error)
	GetNotificationsByUserID(userID string) ([]*Notification, error)
}
