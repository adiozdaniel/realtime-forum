package renders

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"forum/forumapp"
	//"forum/response"
)

// type fort struct {
// 	h *RendersRepo
// }

func TestHomePageHandler(t *testing.T) {
	t.Run("template", func(t *testing.T) {
		// Template cache
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
}
