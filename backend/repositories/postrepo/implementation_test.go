package postrepo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func CreateDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil
	}
	dbString := `
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
		CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);`

	_, err = db.Exec(dbString)
	if err != nil {
		fmt.Println("error creating table")
	}
	return db
}

func TestCreatePost(t *testing.T) {
	db := CreateDb()
	postrepo := &PostRepository{DB: db}
	post := &Post{PostID: "1", PostAuthor: "Makarios", PostTitle: "", PostContent: "", PostCategory: "", UserID: "", HasComments: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), Likes: []*Like{}, Comments: []*Comment{}}
	if _, err := postrepo.CreatePost(post); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}

func TestGetPostByID(t *testing.T) {
}

func TestUpdatePost(t *testing.T) {
}

func TestDeletePost(t *testing.T) {
}

func TestListPosts(t *testing.T) {
}

func TestGetLikesByPostID(t *testing.T) {
}

func TestGetCommentsByPostID(t *testing.T) {
}

func TestGetRepliesByCommentID(t *testing.T) {
}

func TestGetLikesByCommentID(t *testing.T) {
}

func TestGetLikesByReplyID(t *testing.T) {
}

func TestAddLike(t *testing.T) {
}

func TestDislike(t *testing.T) {
}

func TestHasUserLiked(t *testing.T) {
}

func TestCreateComment(t *testing.T) {
}
