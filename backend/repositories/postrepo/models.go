package postrepo

import (
	"forum/forumapp"
	"forum/repositories/shared"
	"time"
)

// Post represents a forum post
type Post struct {
	PostID       string     `json:"post_id"`
	UserID       string     `json:"user_id"`
	PostAuthor   string     `json:"post_author"`
	AuthorImg    string     `json:"author_img,omitempty"`
	PostTitle    string     `json:"post_title"`
	PostContent  string     `json:"post_content"`
	PostImage    string     `json:"post_image,omitempty"`
	PostVideo    string     `json:"post_video,omitempty"`
	PostCategory string     `json:"post_category"`
	HasComments  bool       `json:"post_hasComments"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Likes        []*Like    `json:"likes,omitempty"`
	Comments     []*Comment `json:"comments,omitempty"`
}

// Comment represents a comment in the database
type Comment struct {
	CommentID       string    `json:"comment_id"`
	PostID          string    `json:"post_id"`
	UserID          string    `json:"user_id"`
	Author          string    `json:"user_name"`
	AuthorImg       string    `json:"author_img,omitempty"`
	ParentCommentID string    `json:"parent_comment_id,omitempty"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Likes           []*Like   `json:"likes,omitempty"`
	Replies         []*Reply  `json:"replies,omitempty"`
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
	ReplyID       string    `json:"reply_id"`
	CommentID     string    `json:"comment_id"`
	UserID        string    `json:"user_id"`
	Author        string    `json:"user_name"`
	AuthorImg     string    `json:"author_img,omitempty"`
	ParentReplyID string    `json:"parent_reply_id,omitempty"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Likes         []*Like   `json:"likes,omitempty"`
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
