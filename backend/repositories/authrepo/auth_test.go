package authrepo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
	"forum/repositories/shared"
)

func TestRegisterHandler(t *testing.T) {
	authrepo := &AuthRepo{app: &forumapp.ForumApp{}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}, Sessions: &Sessions{}}
	t.Run("method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/auth/register", nil)
		w := httptest.NewRecorder()
		authrepo.RegisterHandler(w, req)
		if req.Method != http.MethodPost {
			t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestLoginHandler(t *testing.T) {
}

func TestLogoutHandler(t *testing.T) {
}

func TestCheckAuth(t *testing.T) {
}
