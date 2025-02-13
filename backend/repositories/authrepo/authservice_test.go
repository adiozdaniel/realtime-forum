package authrepo

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"forum/repositories/shared"
)

type Db struct {
	Db User
}

func TestRegister(t *testing.T) {
	// db := &Db{Db: User{Email: "tay@gmail.com", Password: "Naaahshshs786$", UserID: "4", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now()}}

	userserv := &UserService{user: &UserRepository{DB: &sql.DB{}}, shared: &shared.SharedConfig{}}
	user := &User{Email: "", Password: "Naaahshshs786$", UserID: "4", UserName: "Abas", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := userserv.Register(user)
	if err.Error() != "email or password cannot be empty" {
		t.Errorf("expected: %v Got %v", errors.New("email or password cannot be empty"), err)
	}
}

func TestLogin(t *testing.T) {
}
