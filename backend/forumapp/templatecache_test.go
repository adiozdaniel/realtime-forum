package forumapp

import (
	//"html/template"
	"testing"
)

func TestCreateTemplatesCache(t *testing.T) {
	h := &TemplateCache{}
	err := h.CreateTemplatesCache()
	if err != nil {
		t.Errorf("Test failed ,%s", err)
	}
}
