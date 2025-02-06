package forumapp

import (
	"testing"
)

func TestCreateTemplatesCache(t *testing.T) {
	h := &TemplateCache{}
	err := h.CreateTemplatesCache()
	if err != nil {
		t.Errorf("Test failed ,%s", err)
	}
}

// func TestGetPage(t *testing.T) {
// 	z := &TemplateCache{}
// 	_, err := z.GetPage("home.page.html")
// 	if err != nil {
// 		t.Errorf("Error getting Page, %s", err)
// 	}
// }
