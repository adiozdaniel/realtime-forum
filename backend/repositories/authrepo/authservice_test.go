package authrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"forum/repositories/shared"
)

type DB struct {
	Db *sql.DB
}

func CreateDb() *DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil
	}
	dbString := `
	CREATE TABLE IF NOT EXISTS users (
		user_id TEXT PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		user_name TEXT,
		image TEXT,
		role TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(dbString)
	if err != nil {
		fmt.Println("error creating table")
	}
	return &DB{Db: db}

}

func TestRegister(t *testing.T) {
	// db := &Db{Db: User{Email: "tay@gmail.com", Password: "Naaahshshs786$", UserID: "4", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	db := CreateDb()

	userserv := &UserService{user: &UserRepository{DB: db.Db}}
	user := &User{Email: "", Password: "Naaahshshs786$", UserID: "4", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := userserv.Register(user)
	if err.Error() != "email or password cannot be empty" {
		t.Errorf("expected: %v Got %v", errors.New("email or password cannot be empty"), err)
	}
}

func TestLogin(t *testing.T) {
	userserv := &UserService{user: &UserRepository{DB: &sql.DB{}}, shared: &shared.SharedConfig{}}
	// user := &User{Email: "", Password: "Naaahshshs786$", UserID: "4", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	_, err := userserv.Login("edu1@gmail.com", "")
	if err.Error() != "email or password cannot be empty" {
		t.Errorf("expected: %v Got %v", errors.New("email or password cannot be empty"), err)
	}
}
