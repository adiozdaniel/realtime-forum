package postrepo

import (
	"database/sql"
	"errors"
	"time"
)

/*
	These concrete implementations of the interfaces interact with the database.
*/

// Post represents a forum post
type Post struct {
	PostID       string    `json:"post_id"`
	UserID       string    `json:"user_id"`
	PostAuthor   string    `json:"post_author"`
	PostTitle    string    `json:"post_title"`
	PostContent  string    `json:"post_content"`
	PostImage    string    `json:"post_image, omitempty"`
	PostVideo    string    `json:"post_video, omitempty"`
	PostCategory string    `json:"post_category"`
	PostLikes    int       `json:"post_likes"`
	HasComments  bool      `json:"post_hasComments`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PostRepository handles database operations for posts
type PostRepository struct {
	DB *sql.DB
}

// NewPostRepository creates a new instance of PostRepository
func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// CreatePost inserts a new post into the database
func (r *PostRepository) CreatePost(post *Post) error {
	query := `INSERT INTO posts (post_id, user_id, post_author, post_title, post_content, post_image, post_video, post_category, post_likes, post_hasComments, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, post.PostID, post.UserID, post.PostAuthor, post.PostTitle, post.PostContent, post.PostImage, post.PostVideo, post.PostCategory, post.PostLikes, post.HasComments, post.CreatedAt, post.UpdatedAt)
	return err
}

// GetPostByID retrieves a post by its ID
func (r *PostRepository) GetPostByID(id string) (*Post, error) {
	query := `SELECT post_id, user_id, post_author, post_title, post_content, post_image, post_video, post_category, post_likes, post_hasComments, created_at, updated_at FROM posts WHERE post_id = ?`
	row := r.DB.QueryRow(query, id)

	post := &Post{}
	err := row.Scan(&post.PostID, &post.UserID, &post.PostAuthor, &post.PostTitle, &post.PostContent, &post.PostImage, &post.PostVideo, &post.PostCategory, &post.PostLikes, &post.HasComments, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return post, nil
}

// UpdatePost updates an existing post in the database
func (r *PostRepository) UpdatePost(post *Post) error {
	query := `UPDATE posts SET post_title = ?, post_content = ?, post_image = ?, post_video = ?, post_category = ?, post_likes = ?, post_hasComments = ?, updated_at = ? WHERE post_id = ?`
	_, err := r.DB.Exec(query, post.PostTitle, post.PostContent, post.PostImage, post.PostVideo, post.PostCategory, post.PostLikes, post.HasComments, post.UpdatedAt, post.PostID)
	return err
}

// DeletePost removes a post from the database by ID
func (r *PostRepository) DeletePost(id string) error {
	query := "DELETE FROM posts WHERE post_id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return errors.New("failed to delete post: " + err.Error())
	}
	return nil
}

// ListPosts retrieves all posts from the database
func (r *PostRepository) ListPosts() ([]*Post, error) {
	query := `SELECT post_id, user_id, post_author, post_title, post_content, post_image, post_video, post_category, post_likes, post_hasComments, created_at, updated_at FROM posts`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.PostID, &post.UserID, &post.PostAuthor, &post.PostTitle, &post.PostContent, &post.PostImage, &post.PostVideo, &post.PostCategory, &post.PostLikes, &post.HasComments, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// AddLike adds a like to a post
func (r *PostRepository) AddLike(id string) error {
	query := `UPDATE posts SET post_likes = post_likes + 1 WHERE post_id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

// DisLike removes a like from a post
func (r *PostRepository) DisLike(id string) error {
	query := `UPDATE posts SET post_likes = post_likes - 1 WHERE post_id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
