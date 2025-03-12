package forumapp

import (
	"html/template"
	"testing"
)

func TestCreateTemplatesCache(t *testing.T) {
	h := &TemplateCache{}
	err := h.CreateTemplatesCache()
	if err != nil {
		t.Errorf("Test failed ,%s", err)
	}
}

func TestGetPage(t *testing.T) {
	// z := newTemplateCache()
	r := make(map[string]*template.Template)
	r["home.page.html"] = template.New("home.page.html")
	z := &TemplateCache{Pages: r}

	_, err := z.GetPage("home.page.html")
	if err != nil {
		t.Errorf("Error getting Page, %s", err)
	}
}
