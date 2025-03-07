package postrepo

import (
	"database/sql"
	"errors"
	"fmt"
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
func (repo *PostRepository) ListPosts() ([]*Post, error) {
	rows, err := repo.DB.Query("SELECT post_id, user_id, post_author, author_img, post_title, post_content, post_image, post_video, post_category, post_hasComments, created_at, updated_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		if err := rows.Scan(&post.PostID, &post.UserID, &post.PostAuthor, &post.AuthorImg, &post.PostTitle, &post.PostContent, &post.PostImage, &post.PostVideo, &post.PostCategory, &post.HasComments, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}

		// Fetch likes
		post.Likes, _ = repo.GetLikesByPostID(post.PostID)
		post.Dislikes, _ = repo.GetDislikesByPostID(post.PostID)

		// Fetch comments
		post.Comments, _ = repo.GetCommentsByPostID(post.PostID)

		// Attach replies to comments
		for _, comment := range post.Comments {
			comment.Replies, _ = repo.GetRepliesByCommentID(comment.CommentID)
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// GetLikesByPostID retrieves all likes for a post by its ID
func (r *PostRepository) GetLikesByPostID(postID string) ([]*Like, error) {
	query := `SELECT like_id, user_id, post_id, comment_id, reply_id FROM likes WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		var commentID, replyID sql.NullString

		err := rows.Scan(&like.LikeID, &like.UserID, &like.PostID, &commentID, &replyID)
		if err != nil {
			return nil, err
		}

		// Convert NULL values to empty strings
		like.CommentID = commentID.String
		like.ReplyID = replyID.String

		if like.CommentID == "" && like.ReplyID == "" {
			likes = append(likes, like)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
}

// GetDislikesByPostID retrieves all dislikes for a post by its ID
func (r *PostRepository) GetDislikesByPostID(postID string) ([]*Like, error) {
	query := `SELECT like_id, user_id, post_id, comment_id, reply_id FROM dislikes WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dislikes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.LikeID, &like.UserID)
		if err != nil {
			return nil, err
		}
		dislikes = append(dislikes, like)
	}
	return dislikes, nil
}

// GetCommentsByPostID retrieves all comments for a post by its ID
func (r *PostRepository) GetCommentsByPostID(postID string) ([]*Comment, error) {
	query := `SELECT comment_id, post_id, post_title, post_author, post_author_img, user_id, user_name, author_img, parent_comment_id, comment, created_at, updated_at FROM comments WHERE post_id = ?`
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		comment := &Comment{}
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.PostTitle, &comment.PostAuthor, &comment.PostAuthorImg, &comment.UserID, &comment.Author, &comment.AuthorImg, &comment.ParentCommentID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Fetch likes for this comment
		comment.Likes, _ = r.GetLikesByCommentID(comment.CommentID)

		// Fetch dislikes for this comment
		comment.Dislikes, _ = r.GetDislikesByCommentID(comment.CommentID)

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

// GetDislikesByCommentID retrieves all dislikes for a comment by its ID
func (r *PostRepository) GetDislikesByCommentID(commentID string) ([]*Like, error) {
	query := `SELECT like_id, user_id FROM dislikes WHERE comment_id = ?`
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

// PostDisLike removes a like from a post
func (r *PostRepository) PostDislike(dislike *Like) (*Like, error) {
	query := `INSERT INTO dislikes (like_id, user_id, post_id, comment_id, reply_id, created_at)
	          VALUES (?, ?, NULLIF(?, ''), NULLIF(?, ''), NULLIF(?, ''), ?)`
	_, err := r.DB.Exec(query, dislike.LikeID, dislike.UserID, dislike.PostID, dislike.CommentID, dislike.ReplyID, dislike.CreatedAt)
	return dislike, err
}

// AddLike adds a like to a post
func (r *PostRepository) AddLike(postlike *Like) (*Like, error) {
	query := `INSERT INTO likes (like_id, user_id, post_id, comment_id, reply_id, created_at)
	          VALUES (?, ?, ?, NULLIF(?, ''), NULLIF(?, ''), ?)`
	_, err := r.DB.Exec(query, postlike.LikeID, postlike.UserID, postlike.PostID, postlike.CommentID, postlike.ReplyID, postlike.CreatedAt)
	return postlike, err
}

// DisLike removes a like from a post
func (r *PostRepository) DisLike(dislike *Like, entityType string) error {
	var query string

	switch entityType {
	case "likes":
		query = `DELETE FROM likes WHERE like_id = ?`
	case "dislikes":
		query = `DELETE FROM dislikes WHERE like_id = ?`
	default:
		return errors.New("invalid entity type")
	}

	_, err := r.DB.Exec(query, dislike.LikeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	return err
}

// HasUserLiked checks if a user has liked a specific post, comment, or reply and returns the like ID if it exists
func (r *PostRepository) HasUserLiked(entityID, userID, entityType string) (string, error) {
	var query string
	var likeID string
	var err error

	switch entityType {
	case "Post":
		query = `SELECT like_id FROM likes WHERE post_id = ? AND user_id = ? AND (comment_id IS NULL OR comment_id = '') AND (reply_id IS NULL OR reply_id = '')`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&likeID)

	case "Comment":
		query = `SELECT like_id FROM likes WHERE comment_id = ? AND user_id = ? AND (reply_id IS NULL OR reply_id = '')`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&likeID)

	case "Reply":
		query = `SELECT like_id FROM likes WHERE reply_id = ? AND user_id = ?`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&likeID)

	default:
		return "", errors.New("invalid entity type")
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // No like found, return empty string
		}
		return "", err
	}

	return likeID, nil
}

// HasUserDisliked checks if a user has disliked a specific post, comment, or reply and returns the dislike ID if it exists
func (r *PostRepository) HasUserDisliked(entityID, userID, entityType string) (string, error) {
	var query string
	var dislikeID string
	var err error

	switch entityType {
	case "Post":
		query = `SELECT like_id FROM dislikes WHERE post_id = ? AND user_id = ? AND (comment_id IS NULL OR comment_id = '') AND (reply_id IS NULL OR reply_id = '')`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&dislikeID)

	case "Comment":
		query = `SELECT like_id FROM dislikes WHERE comment_id = ? AND user_id = ? AND (reply_id IS NULL OR reply_id = '')`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&dislikeID)

	case "Reply":
		query = `SELECT like_id FROM dislikes WHERE reply_id = ? AND user_id = ?`
		err = r.DB.QueryRow(query, entityID, userID).Scan(&dislikeID)

	default:
		return "", errors.New("invalid entity type")
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // No dislike found, return empty string
		}
		return "", err
	}

	return dislikeID, nil
}

// Comments
// CreateComment creates a new comment
func (r *PostRepository) CreateComment(comment *Comment) (*Comment, error) {
	query := `INSERT INTO comments (comment_id, post_id, post_title, post_author, post_author_img, user_id, user_name, author_img, parent_comment_id, comment, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?, NULLIF(?, ''), ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.CommentID, comment.PostID, comment.PostTitle, comment.PostAuthor, comment.PostAuthorImg, comment.UserID, comment.Author, comment.AuthorImg, comment.ParentCommentID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	return comment, err
}

// UpdateComment updates an existing comment
func (r *PostRepository) UpdateComment(comment *Comment) (*Comment, error) {
	query := `UPDATE comments SET post_title = ?, post_author = ?, post_author_img = ?, user_id = ?, user_name = ?, author_img = ?, parent_comment_id = ?, comment = ?, updated_at = ? WHERE comment_id = ?`
	_, err := r.DB.Exec(query, comment.PostTitle, comment.PostAuthor, comment.PostAuthorImg, comment.UserID, comment.Author, comment.AuthorImg, comment.ParentCommentID, comment.Content, comment.UpdatedAt, comment.CommentID)
	return comment, err
}

// DeleteComment deletes a comment
func (r *PostRepository) DeleteComment(id string) error {
	query := `DELETE FROM comments WHERE comment_id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

// CreateReply creates a new reply
func (r *PostRepository) CreateReply(reply *Reply) (*Reply, error) {
	query := `INSERT INTO replies (reply_id, comment_id, user_id, user_name, author_img, parent_reply_id, content, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, NULLIF(?, ''), ?, ?, ?)`
	_, err := r.DB.Exec(query, reply.ReplyID, reply.CommentID, reply.UserID, reply.Author, reply.AuthorImg, reply.ParentReplyID, reply.Content, reply.CreatedAt, reply.UpdatedAt)
	return reply, err
}

// GetPostsByUserID retrieves all posts created by a specific user
func (r *PostRepository) GetPostsByUserID(userID string) ([]*Post, error) {
	query := `SELECT post_id, user_id, post_author, author_img, post_title, post_content, post_image, post_video, post_category, post_hasComments, created_at, updated_at FROM posts WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
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

		post.Likes, _ = r.GetLikesByPostID(post.PostID)
		post.Dislikes, _ = r.GetDislikesByPostID(post.PostID)
		post.Comments, _ = r.GetCommentsByPostID(post.PostID)
		posts = append(posts, post)
	}
	return posts, nil
}

// GetLikedPostsByUserID retrieves all posts liked by a user
func (r *PostRepository) GetLikedPostsByUserID(userID string) ([]*Post, error) {
	likes, err := r.GetLikesByUserID(userID)
	if err != nil {
		return nil, err
	}

	var posts []*Post

	for _, like := range likes {
		if like.CommentID == "" && like.ReplyID == "" {
			post, err := r.GetPostByLikeID(like.LikeID, userID)
			if post == nil || err != nil {
				continue
			}

			post.Likes, _ = r.GetLikesByPostID(post.PostID)
			post.Dislikes, _ = r.GetDislikesByPostID(post.PostID)
			post.Comments, _ = r.GetCommentsByPostID(post.PostID)

			posts = append(posts, post)
		}
	}

	return posts, nil
}

// GetCommentsByUserID retrieves all comments created by a specific user
func (r *PostRepository) GetCommentsByUserID(userID string) ([]*Comment, error) {
	query := `SELECT comment_id, post_id, post_title, post_author, post_author_img, user_id, user_name, author_img, parent_comment_id, comment, created_at, updated_at FROM comments WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		comment := &Comment{}
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.PostTitle, &comment.PostAuthor, &comment.PostAuthorImg, &comment.UserID, &comment.Author, &comment.AuthorImg, &comment.ParentCommentID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// GetRepliesByUserID retrieves all replies created by a specific user
func (r *PostRepository) GetRepliesByUserID(userID string) ([]*Reply, error) {
	query := `SELECT reply_id, comment_id, user_id, user_name, author_img, parent_reply_id, content, created_at, updated_at FROM replies WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
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
		replies = append(replies, reply)
	}
	return replies, nil
}

// GetLikesByUserID retrieves all likes created by a specific user
func (r *PostRepository) GetLikesByUserID(userID string) ([]*Like, error) {
	query := `SELECT like_id, user_id, post_id, comment_id, reply_id FROM likes WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		var commentID, replyID sql.NullString

		err := rows.Scan(&like.LikeID, &like.UserID, &like.PostID, &commentID, &replyID)
		if err != nil {
			return nil, err
		}

		// Convert NULL values to empty strings
		like.CommentID = commentID.String
		like.ReplyID = replyID.String

		likes = append(likes, like)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
}

// GetDislikesByUserID retrieves all dislikes created by a specific user
func (r *PostRepository) GetDislikesByUserID(userID string) ([]*Like, error) {
	query := `SELECT like_id, user_id FROM dislikes WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dislikes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.LikeID, &like.UserID)
		if err != nil {
			return nil, err
		}
		dislikes = append(dislikes, like)
	}
	return dislikes, nil
}

// AddActivity adds a new activity to the database
func (r *PostRepository) AddActivity(activity *Activity) (*Activity, error) {
	query := `INSERT INTO activities (activity_id, user_id, activity_type, activity_data, created_at)
	           VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, activity.ActivityID, activity.UserId, activity.ActivityType, activity.ActivityData, activity.CreatedAt)
	return activity, err
}

// GetPostByLikeID retrieves a post by its like ID and user ID
func (r *PostRepository) GetPostByLikeID(likeID, userID string) (*Post, error) {
	query := `SELECT post_id, user_id, post_author, author_img, post_title, post_content, 
	                 post_image, post_video, post_category, post_hasComments, created_at, updated_at
	          FROM posts 
	          WHERE post_id = (SELECT post_id FROM likes WHERE like_id = ? AND user_id = ?)`

	var post Post
	var authorImg, postImage, postVideo sql.NullString // Handle NULL values

	err := r.DB.QueryRow(query, likeID, userID).Scan(
		&post.PostID, &post.UserID, &post.PostAuthor, &authorImg,
		&post.PostTitle, &post.PostContent, &postImage, &postVideo,
		&post.PostCategory, &post.HasComments, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Convert NULL values to empty strings
	post.AuthorImg = authorImg.String
	post.PostImage = postImage.String
	post.PostVideo = postVideo.String

	return &post, nil
}

// GetCommentByLikeID retrieves a comment by its like ID and user ID
func (r *PostRepository) GetCommentByLikeID(likeID, userID string) (*Comment, error) {
	query := `SELECT comment_id, user_id, post_id, post_title, post_author, post_author_img, comment_content, comment_likes, created_at, updated_at 
	          FROM comments 
	          WHERE comment_id IN (SELECT comment_id FROM likes WHERE like_id = ? AND user_id = ?) AND user_id = ?`

	var comment Comment
	err := r.DB.QueryRow(query, likeID, userID, userID).Scan(
		&comment.CommentID, &comment.UserID, &comment.PostID, &comment.PostTitle, &comment.PostAuthor, &comment.PostAuthorImg, &comment.Content,
		&comment.Likes, &comment.CreatedAt, &comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetActivitiesByUserID retrieves all activities created by a specific user
func (r *PostRepository) GetActivitiesByUserID(userID string) ([]*Activity, error) {
	query := `SELECT activity_id, user_id, activity_type, activity_data, created_at 
	          FROM activities 
	          WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*Activity
	for rows.Next() {
		activity := &Activity{}
		err := rows.Scan(&activity.ActivityID, &activity.UserId, &activity.ActivityType, &activity.ActivityData, &activity.CreatedAt)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

// GetCommentByID retrieves a comment by its ID
func (r *PostRepository) GetCommentByID(id string) (*Comment, error) {
	query := `SELECT comment_id, post_id, post_title, post_author, post_author_img, user_id, user_name, author_img, parent_comment_id, comment, created_at, updated_at 
	          FROM comments 
	          WHERE comment_id = ?`
	var comment Comment
	err := r.DB.QueryRow(query, id).Scan(
		&comment.CommentID, &comment.PostID, &comment.PostTitle, &comment.PostAuthor, &comment.PostAuthorImg, &comment.UserID, &comment.Author, &comment.AuthorImg, &comment.ParentCommentID, &comment.Content,
		&comment.CreatedAt, &comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// CreateNotification creates a new notification
func (r *PostRepository) CreateNotification(n *Notification) (*Notification, error) {
	query := `INSERT INTO notifications (notification_id, user_id, sender_id, post_id, comment_id, reply_id, like_id, dislike_id, notification_type, message, is_read, created_at) 
	          VALUES (?, ?, ?, NULLIF(?, ''), NULLIF(?, ''), NULLIF(?, ''), NULLIF(?, ''), NULLIF(?, ''), ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, n.NotificationID, n.UserId, n.SenderID, n.PostID, n.CommentID, n.ReplyID, n.LikeID, n.DislikeID, n.NotificationType, n.Message, n.IsRead, n.CreatedAt)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// GetNotificationsByUserID retrieves all notifications for a user by their ID
func (r *PostRepository) GetNotificationsByUserID(userID string) ([]*Notification, error) {
	query := `SELECT notification_id, user_id, sender_id, post_id, comment_id, reply_id, like_id, dislike_id, notification_type, message, is_read, created_at 
	          FROM notifications 
	          WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*Notification
	for rows.Next() {
		notification := &Notification{}
		err := rows.Scan(&notification.NotificationID, &notification.UserId, &notification.SenderID, &notification.PostID, &notification.CommentID, &notification.ReplyID, &notification.LikeID, &notification.DislikeID, &notification.NotificationType, &notification.Message, &notification.IsRead, &notification.CreatedAt)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
