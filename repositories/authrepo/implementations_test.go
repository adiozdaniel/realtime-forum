package authrepo

import (
	//"database/sql"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	id    = rand.Intn(100)
	idstr = strconv.Itoa(id)
)

func TestCreateUser(t *testing.T) {
	db := CreateDb()
	userrepo := &UserRepository{DB: db.Db}

	user := &User{Email: "h" + idstr + "@R.COM", Password: "Naaahshshs786$", UserID: idstr, UserName: "Abasa", Bio: "Wanderer", CreatedAt: time.Now(), UpdatedAt: time.Now().Add(1 * time.Hour)}
	err := userrepo.CreateUser(user)
	if err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db := CreateDb()
	userrepo := &UserRepository{DB: db.Db}
	_, err := userrepo.GetUserByEmail("h" + idstr + "@R.COM")
	if err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}

func TestDeleteUser(t *testing.T) {
	db := CreateDb()
	userrepo := &UserRepository{DB: db.Db}
	err := userrepo.DeleteUser("h" + idstr + "@R.COM")
	if err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}

func TestGetUserByID(t *testing.T) {
	db := CreateDb()
	user := &User{Email: "h@R.COM", Password: "Naaahshshs786$", UserID: "2", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now().Add(1 * time.Hour)}
	userrepo := &UserRepository{DB: db.Db}
	_, err := userrepo.GetUserByID(user)
	if err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}

func TestUpdateUser(t *testing.T) {
	db := CreateDb()
	userrepo := &UserRepository{DB: db.Db}
	user := &User{Email: "nash@com", Password: "Naaahshshs786$", UserID: "1", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now().Add(1 * time.Hour)}
	_, err := userrepo.UpdateUser(user)
	if err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}
