package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
	"forum/response"
)

func TestRegisterHandler(t *testing.T) {
	// Template cache
	r := make(map[string]*template.Template)
	r["home.page.html"] = template.New("home.page.html")
	tmplcach := &forumapp.TemplateCache{Pages: r}

	fapp := &forumapp.ForumApp{}
	fapp.Tmpls = tmplcach
	jsonres := &response.JSONRes{}

	h := &Repo{app: fapp, res: jsonres}
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", nil)
	w := httptest.NewRecorder()
	h.RegisterHandler(w, req)
	if req.Method != http.MethodPost {
		t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	// Template cache
	r := make(map[string]*template.Template)
	r["home.page.html"] = template.New("home.page.html")
	tmplcach := &forumapp.TemplateCache{Pages: r}

	fapp := &forumapp.ForumApp{}
	fapp.Tmpls = tmplcach
	jsonres := &response.JSONRes{}
	h := &Repo{app: fapp, res: jsonres}
	req := httptest.NewRequest(http.MethodGet, "/api/auth/login", nil)
	w := httptest.NewRecorder()
	h.LoginHandler(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestPostsHandler(t *testing.T) {
	// Template cache
	r := make(map[string]*template.Template)
	r["home.page.html"] = template.New("home.page.html")
	tmplcach := &forumapp.TemplateCache{Pages: r}

	fapp := &forumapp.ForumApp{}
	fapp.Tmpls = tmplcach
	jsonres := &response.JSONRes{}
	h := &Repo{app: fapp, res: jsonres}
	req := httptest.NewRequest(http.MethodGet, "/api/auth/posts", nil)
	w := httptest.NewRecorder()
	h.PostsHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
