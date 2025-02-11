package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
	"forum/response"
)

type fort struct {
	h *Repo
}

func TestHomePageHandler(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		// tmplcach := forumapp.NewTemplateCache()
		// tmplcach.CreateTemplatesCache()
		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &Repo{app: fapp}
		// h := &Repo{}
		// h.app.Tmpls.CreateTemplatesCache() // Pages["home.page.html"]=
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		writer := httptest.NewRecorder()
		// fmt.Println(req, writer, h)
		h.HomePageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
	t.Run("Test httpmethod", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &Repo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		writer := httptest.NewRecorder()

		h.HomePageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
}

func TestLoginPage(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &Repo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", nil)
		writer := httptest.NewRecorder()

		h.LoginPage(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("template", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		jsonres := &response.JSONRes{}
		h := &Repo{app: fapp, res: jsonres}

		req := httptest.NewRequest(http.MethodGet, "/api/auth/login", nil)
		writer := httptest.NewRecorder()

		h.LoginPage(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}

func TestSignUpPage(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &Repo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/auth-sign-up", nil)
		writer := httptest.NewRecorder()

		h.SignUpPage(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("template", func(t *testing.T) {
		// Template cache
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		jsonres := &response.JSONRes{}
		h := &Repo{app: fapp, res: jsonres}

		req := httptest.NewRequest(http.MethodGet, "/auth/sign-up", nil)
		writer := httptest.NewRecorder()

		h.SignUpPage(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}
