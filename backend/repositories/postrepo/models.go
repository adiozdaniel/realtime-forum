package postrepo

import (
	"database/sql"
	"forum/forumapp"
	"forum/repositories/shared"
	"time"
)

// Post represents a forum post
type Post struct {
	PostID       string     `json:"post_id"`
	UserID       string     `json:"user_id"`
	PostAuthor   string     `json:"post_author"`
	AuthorImg    string     `json:"author_img"`
	PostTitle    string     `json:"post_title"`
	PostContent  string     `json:"post_content"`
	PostImage    string     `json:"post_image,omitempty"`
	PostVideo    string     `json:"post_video,omitempty"`
	PostCategory string     `json:"post_category"`
	HasComments  bool       `json:"post_hasComments"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Likes        []*Like    `json:"likes"`
	Dislikes     []*Like    `json:"dislikes"`
	Comments     []*Comment `json:"comments"`
}

// Comment represents a comment in the database
type Comment struct {
	CommentID       string         `json:"comment_id"`
	PostID          string         `json:"post_id"`
	UserID          string         `json:"user_id"`
	Author          string         `json:"user_name"`
	AuthorImg       string         `json:"author_img"`
	ParentCommentID sql.NullString `json:"parent_comment_id,omitempty"`
	Content         string         `json:"comment"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Likes           []*Like        `json:"likes"`
	Replies         []*Reply       `json:"replies"`
}

// Like represents likes
type Like struct {
	LikeID    string    `json:"like_id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id,omitempty"`
	CommentID string    `json:"comment_id,omitempty"`
	ReplyID   string    `json:"reply_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Replies represents a comment reply
type Reply struct {
	ReplyID       string         `json:"reply_id"`
	CommentID     string         `json:"comment_id"`
	UserID        string         `json:"user_id"`
	Author        string         `json:"user_name"`
	AuthorImg     string         `json:"author_img"`
	ParentReplyID sql.NullString `json:"parent_reply_id,omitempty"`
	Content       string         `json:"content"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Likes         []*Like        `json:"likes,omitempty"`
}

// Activity represents user activities
type Activity struct {
	ActivityID   string    `json:"activity_id"`
	UserId       string    `json:"user_id"`
	ActivityType string    `json:"activity_type"`
	ActivityData string    `json:"activity_data"`
	CreatedAt    time.Time `json:"created_at"`
}

// Notification represents a notification
type Notification struct {
	NotificationID   string    `json:"notification_id"`
	UserId           string    `json:"user_id"`
	SenderID         string    `json:"sender_id"`
	PostID           string    `json:"post_id,omitempty"`
	CommentID        string    `json:"comment_id,omitempty"`
	ReplyID          string    `json:"reply_id,omitempty"`
	LikeID           string    `json:"like_id,omitempty"`
	DislikeID        string    `json:"dislike_id,omitempty"`
	NotificationType string    `json:"notification_type"`
	Message          string    `json:"message"`
	IsRead           bool      `json:"is_read"`
	CreatedAt        time.Time `json:"created_at"`
}

// PostRepo represents posts repository
type PostsRepo struct {
	app    *forumapp.ForumApp
	res    *shared.JSONRes
	post   *PostService
	shared *shared.SharedConfig
}

// NewPost returns a new Post instance
func NewPostsRepo(app *forumapp.ForumApp) *PostsRepo {
	res := shared.NewJSONRes()
	shared := shared.NewSharedConfig()
	postRepo := NewPostRepository(app.Db.Query)
	postService := NewPostService(postRepo)

	return &PostsRepo{
		app:    app,
		res:    res,
		post:   postService,
		shared: shared,
	}
}
