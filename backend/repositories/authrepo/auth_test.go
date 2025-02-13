package authrepo

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
	"forum/repositories/shared"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		authrepo := &AuthRepo{app: &forumapp.ForumApp{}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}, Sessions: &Sessions{}}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/register", nil)
		w := httptest.NewRecorder()
		authrepo.RegisterHandler(w, req)
		if req.Method != http.MethodPost {
			t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
	t.Run("data incomplete or missing", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		incomplete := []byte(`{"username": "John Doe"}`)

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}, Sessions: &Sessions{}}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(incomplete))
		w := httptest.NewRecorder()
		authrepo.RegisterHandler(w, req)
		if w.Code != http.StatusConflict {
			t.Errorf("expected %d got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestLoginHandler(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		authrepo := &AuthRepo{app: &forumapp.ForumApp{}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}, Sessions: &Sessions{}}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", nil)
		w := httptest.NewRecorder()
		authrepo.LoginHandler(w, req)
		if req.Method != http.MethodPost {
			t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestLogoutHandler(t *testing.T) {
}

func TestCheckAuth(t *testing.T) {
}
