package postrepo

import (
	"database/sql"
	"errors"
)

/*
	These concrete implementations of the interfaces interact with the database.
*/

// PostRepository handles database operations for posts
type PostRepository struct {
	DB *sql.DB
}

// NewPostRepository creates a new instance of PostRepository
func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

// CreatePost inserts a new post into the database
func (r *PostRepository) CreatePost(post *Post) (*Post, error) {
	query := `INSERT INTO posts (post_id, user_id, post_author, author_img, post_title, post_content, post_image, post_video, post_category, post_hasComments, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, post.PostID, post.UserID, post.PostAuthor, post.AuthorImg, post.PostTitle, post.PostContent, post.PostImage, post.PostVideo, post.PostCategory, post.HasComments, post.CreatedAt, post.UpdatedAt)
	return post, err
}

// GetPostByID retrieves a post by its ID
func (r *PostRepository) GetPostByID(id string) (*Post, error) {
	query := `SELECT post_id, user_id, post_author, author_img, post_title, post_content, post_image, post_video, post_category, post_hasComments, created_at, updated_at FROM posts WHERE post_id = ?`
	row := r.DB.QueryRow(query, id)

	post := &Post{}
	err := row.Scan(&post.PostID, &post.UserID, &post.PostAuthor, &post.AuthorImg, &post.PostTitle, &post.PostContent, &post.PostImage, &post.PostVideo, &post.PostCategory, &post.HasComments, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return post, nil
}

// UpdatePost updates an existing post in the database
func (r *PostRepository) UpdatePost(post *Post) (*Post, error) {
	query := `UPDATE posts SET post_title = ?, post_content = ?, post_image = ?, post_video = ?, post_category = ?, post_hasComments = ?, updated_at = ? WHERE post_id = ?`
	_, err := r.DB.Exec(query, post.PostTitle, post.PostContent, post.PostImage, post.PostVideo, post.PostCategory, post.HasComments, post.UpdatedAt, post.PostID)
	return post, err
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
	query := `SELECT post_id, user_id, post_author, author_img, post_title, post_content, post_image, post_video, post_category, post_hasComments, created_at, updated_at FROM posts`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.PostID, &post.UserID, &post.PostAuthor, &post.AuthorImg, &post.PostTitle, &post.PostContent, &post.PostImage, &post.PostVideo, &post.PostCategory, &post.HasComments, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Fetch likes and comments
		post.Likes, _ = r.GetLikesByPostID(post.PostID)
		post.Comments, _ = r.GetCommentsByPostID(post.PostID)

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetLikesByPostID retrieves all likes for a post by its ID
func (r *PostRepository) GetLikesByPostID(postID string) ([]*Like, error) {
	query := `SELECT like_id, user_id FROM likes WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.LikeID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// GetCommentsByPostID retrieves all comments for a post by its ID
func (r *PostRepository) GetCommentsByPostID(postID string) ([]*Comment, error) {
	query := `SELECT comment_id, user_id, user_name, author_img, parent_comment_id, content, created_at, updated_at FROM comments WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		comment := &Comment{}
		err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.Author, &comment.AuthorImg, &comment.ParentCommentID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Fetch likes for this comment
		comment.Likes, _ = r.GetLikesByCommentID(comment.CommentID)

		// Fetch replies for this comment
		comment.Replies, _ = r.GetRepliesByCommentID(comment.CommentID)

		comments = append(comments, comment)
	}

	return comments, nil
}

// GetRepliesByCommentID retrieves all replies for a comment by its ID
func (r *PostRepository) GetRepliesByCommentID(commentID string) ([]*Reply, error) {
	query := `SELECT reply_id, comment_id, user_id, user_name, author_img, parent_reply_id, content, created_at, updated_at FROM replies WHERE comment_id = ?`
	rows, err := r.DB.Query(query, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []*Reply
	for rows.Next() {
		reply := &Reply{}
		err := rows.Scan(&reply.ReplyID, &reply.CommentID, &reply.UserID, &reply.Author, &reply.AuthorImg, &reply.ParentReplyID, &reply.Content, &reply.CreatedAt, &reply.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Fetch likes for this reply
		reply.Likes, _ = r.GetLikesByReplyID(reply.ReplyID)

		replies = append(replies, reply)
	}
	return replies, nil
}

// GetLikesByCommentID retrieves all likes for a comment by its ID
func (r *PostRepository) GetLikesByCommentID(commentID string) ([]*Like, error) {
	query := `SELECT like_id, user_id FROM likes WHERE comment_id = ?`
	rows, err := r.DB.Query(query, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.LikeID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// GetLikesByReplyID retrieves all likes for a reply by its ID
func (r *PostRepository) GetLikesByReplyID(replyID string) ([]*Like, error) {
	query := `SELECT like_id, user_id FROM likes WHERE reply_id = ?`
	rows, err := r.DB.Query(query, replyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.LikeID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

// AddLike adds a like to a post
func (r *PostRepository) AddLike(postlike *Like) (*Like, error) {
	query := `INSERT INTO likes (like_id, user_id, post_id, comment_id, reply_id, created_at)
	          VALUES (?, ?, ?, NULLIF(?, ''), NULLIF(?, ''), ?)`
	_, err := r.DB.Exec(query, postlike.LikeID, postlike.UserID, postlike.PostID, postlike.CommentID, postlike.ReplyID, postlike.CreatedAt)
	return postlike, err
}

// DisLike removes a like from a post
func (r *PostRepository) DisLike(postdislike *Like) error {
	query := `DELETE FROM likes WHERE like_id = ?`
	_, err := r.DB.Exec(query, postdislike.LikeID)
	return err
}

// HasUserLiked checks if a user has liked a specific post, comment, or reply and returns the like ID if it exists
func (r *PostRepository) HasUserLiked(entityID, userID, entityType string) (string, error) {
	var query string
	var likeID string

	switch entityType {
	case "Post":
		query = `SELECT like_id FROM likes WHERE post_id = ? AND user_id = ?`
	case "Comment":
		query = `SELECT like_id FROM likes WHERE comment_id = ? AND user_id = ?`
	case "Reply":
		query = `SELECT like_id FROM likes WHERE reply_id = ? AND user_id = ?`
	default:
		return "", errors.New("invalid entity type")
	}

	err := r.DB.QueryRow(query, entityID, userID).Scan(&likeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // No like found, return empty string
		}
		return "", err
	}

	return likeID, nil
}
