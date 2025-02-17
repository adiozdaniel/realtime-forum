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
}

func TestGetLikesByCommentID(t *testing.T) {
}

func TestGetLikesByReplyID(t *testing.T) {
}

func TestAddLike(t *testing.T) {
}

func TestDislike(t *testing.T) {
}

// func TestHasUserLiked(t *testing.T) {//responsive just missing table
// 	db := CreateDb()
// 	postrepo := &PostRepository{DB: db}

// 	if _, err := postrepo.HasUserLiked("1", "", "Post"); err != nil {
// 		t.Errorf("expected %v, got %v", nil, err)
// 	}
// }

func TestCreateComment(t *testing.T) {
}
