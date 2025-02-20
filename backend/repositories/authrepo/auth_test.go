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
		authrepo := &AuthRepo{app: &forumapp.ForumApp{}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}
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

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}

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
		authrepo := &AuthRepo{app: &forumapp.ForumApp{}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}
		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", nil)
		w := httptest.NewRecorder()
		authrepo.LoginHandler(w, req)
		if req.Method != http.MethodPost {
			t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
	t.Run("incorrrect or malformed inputs", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		data := []byte(`{"username": "John Doe"}`)

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		authrepo.LoginHandler(w, req)
		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected %d got %d", http.StatusUnauthorized, w.Code)
		}
	})
}

func TestLogoutHandler(t *testing.T) {
	t.Run("missing session cookie", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		data := []byte(`{"username": "John Doe"}`)

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/logout", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		authrepo.LogoutHandler(w, req)
		if _, err := req.Cookie("session_cookie"); err != http.ErrNoCookie {
			t.Errorf("expected %d got %d", http.ErrNoCookie, w.Code)
		}
	})
	t.Run("method", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		data := []byte(`{"username": "John Doe"}`)

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}

		req := httptest.NewRequest(http.MethodGet, "/api/auth/logout", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		authrepo.LogoutHandler(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestCheckAuth(t *testing.T) {
	t.Run("missing session cookie", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		data := []byte(`{"username": "John Doe"}`)

		authrepo := &AuthRepo{app: &forumapp.ForumApp{Tmpls: tmplcach}, res: &shared.JSONRes{}, user: &UserService{}, shared: &shared.SharedConfig{}}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/check", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		authrepo.CheckAuth(w, req)
		if _, err := req.Cookie("session_cookie"); err != http.ErrNoCookie {
			t.Errorf("expected %d got %d", http.ErrNoCookie, w.Code)
		}
	})
}
