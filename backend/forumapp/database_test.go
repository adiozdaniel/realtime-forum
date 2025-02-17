package forumapp

import (
	"database/sql"
	"testing"
)

func TestCreateTables(t *testing.T) {
	db, _ := sql.Open("sqlite3", "./forum.db")
	tb := &TableManager{db: db}
	if err := tb.CreateTables(); err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
}
