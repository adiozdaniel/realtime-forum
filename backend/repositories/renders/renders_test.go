package renders

import (
	"net/http/httptest"
	"testing"

	"forum/forumapp"
)

// actually this is duplication as the page handlers call this hence test it anyways
func TestRenderTemplate(t *testing.T) {
	tempcache := &forumapp.TemplateCache{}
	app := &forumapp.ForumApp{Tmpls: tempcache}
	repo := &RendersRepo{app: app}
	err := repo.RenderTemplate(httptest.NewRecorder(), "home45.page.html", "sasa")
	if err == nil {
		t.Fatal(err)
	}
}
