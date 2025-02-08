package repositories

import (
	"forum/forumapp"
	"forum/response"
	"time"
)

type Repo struct {
	app     *forumapp.ForumApp
	res     *response.JSONRes
	post    *PostService
	comment *CommentService
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	// get the initialised db connection
	db := app.Db.Query

	// initialise services
	post := NewPostService(db)
	comment := NewCommentService(db)
	return &Repo{
		app, response.NewJSONRes(),
		post, comment,
	}
}

// Post represents a post in the database
type Post struct {
	PostID       string    `json:"post_id"`
	UserID       string    `json:"user_id"`
	PostTitle    string    `json:"post_title"`
	PostContent  string    `json:"post_content"`
	PostImage    string    `json:"post_image,omitempty"`
	PostVideo    string    `json:"post_video,omitempty"`
	PostCategory string    `json:"post_category"`
	PostLikes    int       `json:"post_likes"`
	PostDislikes int       `json:"post_dislikes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Comment represents a comment in the database
type Comment struct {
	CommentID       string    `json:"comment_id"`
	PostID          string    `json:"post_id"`
	UserID          string    `json:"user_id"`
	ParentCommentID string    `json:"parent_comment_id,omitempty"`
	Comment         string    `json:"comment"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
