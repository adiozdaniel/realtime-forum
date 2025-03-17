package postrepo

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	id    = rand.Intn(100)
	idstr = strconv.Itoa(id)
)

func CreateDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil
	}
	tables := map[string]string{
		"users": `
		CREATE TABLE IF NOT EXISTS users (
			user_id TEXT PRIMARY KEY,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			user_name TEXT,
			image TEXT,
			role TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);`,

		"posts": `
		CREATE TABLE IF NOT EXISTS posts (
			post_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			post_author TEXT NOT NULL,
			author_img TEXT,
			post_title TEXT NOT NULL,
			post_content TEXT NOT NULL,
			post_image TEXT,
			post_video TEXT,
			post_category TEXT NOT NULL,
			post_hasComments BOOLEAN DEFAULT TRUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_posts_post_id ON posts(post_id);
		CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);`,

		"comments": `
		CREATE TABLE IF NOT EXISTS comments (
			comment_id TEXT PRIMARY KEY,
			post_id TEXT NOT NULL,
			post_title TEXT,
			post_author TEXT,
			post_author_img TEXT,
			user_id TEXT NOT NULL,
			user_name TEXT NOT NULL,
			author_img TEXT,
			parent_comment_id TEXT,
			comment TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (parent_comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE
		);CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
		CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);`,

		"likes": `
		CREATE TABLE IF NOT EXISTS likes (
			like_id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			post_id TEXT,
			comment_id TEXT,
			reply_id TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE
		);
		CREATE INDEX IF NOT EXISTS idx_likes_user_id ON likes(user_id);`,

		"replies": `
		CREATE TABLE IF NOT EXISTS replies (
			reply_id TEXT PRIMARY KEY,
			comment_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			user_name TEXT NOT NULL,
			author_img TEXT NOT NULL,
			parent_reply_id TEXT,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (parent_reply_id) REFERENCES replies(reply_id) ON DELETE CASCADE
		);`,
	}
	for _, dbString := range tables {
		_, err = db.Exec(dbString)
		if err != nil {
			fmt.Println("error creating table")
		}

	}
	return db
}

func TestCreatePost(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}
	post := &Post{PostID: "", PostAuthor: "Makarios", PostTitle: "", PostContent: "", PostCategory: "", UserID: "", HasComments: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Likes: []*Like{}, Comments: []*Comment{}}
	if _, err := postrepo.CreatePost(post); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestGetPostByID(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.GetPostByID(""); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestUpdatePost(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}
	post := &Post{PostID: "", PostAuthor: "", PostTitle: "", PostContent: "AAH", PostCategory: "", UserID: "", HasComments: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Likes: []*Like{}, Comments: []*Comment{}}
	if _, err := postrepo.UpdatePost(post); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestDeletePost(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if err := postrepo.DeletePost(""); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestListPosts(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.ListPosts(); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

// func TestGetLikesByPostID(t *testing.T) {//not reponsive interms of %
// 	db := CreateDb()
// 	postrepo := &PostRepository{DB: db}

// 	if _, err := postrepo.GetLikesByPostID(""); err != nil {
// 		t.Errorf("expected %v, got %v", nil, err)
// 	}
// }

// func TestGetCommentsByPostID(t *testing.T) {//not responsive interms of %
// 	db := CreateDb()
// 	postrepo := &PostRepository{DB: db}

// 	if _, err := postrepo.GetCommentsByPostID("1"); err != nil {
// 		t.Errorf("expected %v, got %v", nil, err)
// 	}
// }

func TestGetRepliesByCommentID(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.GetRepliesByCommentID("1"); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestGetLikesByCommentID(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.GetLikesByCommentID("1"); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestGetLikesByReplyID(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.GetLikesByReplyID("1"); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestAddLike(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}
	like := &Like{LikeID: idstr}

	if _, err := postrepo.AddLike(like); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestDislike(t *testing.T) {
}

func TestHasUserLiked(t *testing.T) { // responsive just missing table
	db := CreateDb()
	postrepo := &PostRepository{DB: db}

	if _, err := postrepo.HasUserLiked("1", "", "Post"); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestCreateComment(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}
	comment := &Comment{CommentID: idstr}

	if _, err := postrepo.CreateComment(comment); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}
