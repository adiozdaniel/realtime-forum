package renders

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
)

func TestHomePageHandler(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.ipage.html"] = template.New("home.ipage.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		writer := httptest.NewRecorder()
		h.HomePageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
	t.Run("Method", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}
		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		writer := httptest.NewRecorder()
		h.HomePageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
}

func TestLoginPageHandler(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/api/auth/login", nil)
		writer := httptest.NewRecorder()

		h.LoginPageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("template", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodGet, "/api/auth/login", nil)
		writer := httptest.NewRecorder()

		h.LoginPageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}

func TestSignUpPageHandler(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/auth-sign-up", nil)
		writer := httptest.NewRecorder()

		h.SignUpPageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("template", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach

		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodGet, "/auth/sign-up", nil)
		writer := httptest.NewRecorder()

		h.SignUpPageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}

func TestModeratorPageHandler(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/moderator", nil)
		writer := httptest.NewRecorder()

		h.ModeratorPageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("Template", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodGet, "/moderator", nil)
		writer := httptest.NewRecorder()

		h.ModeratorPageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}

func TestAdminPageHandler(t *testing.T) {
	t.Run("Test httpmethod", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodPost, "/admin", nil)
		writer := httptest.NewRecorder()

		h.AdminPageHandler(writer, req)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Expected Method %d,got %d", http.StatusForbidden, writer.Code)
		}
	})
	t.Run("Template", func(t *testing.T) {
		r := make(map[string]*template.Template)
		r["home.page.html"] = template.New("home.page.html")
		tmplcach := &forumapp.TemplateCache{Pages: r}

		fapp := &forumapp.ForumApp{}
		fapp.Tmpls = tmplcach
		h := &RendersRepo{app: fapp}

		req := httptest.NewRequest(http.MethodGet, "/admin", nil)
		writer := httptest.NewRecorder()

		h.AdminPageHandler(writer, req)
		if writer.Code != http.StatusInternalServerError {
			t.Errorf("Expected Method %d,got %d", http.StatusInternalServerError, writer.Code)
		}
	})
}
