package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	//"forum/forumapp"
)

type fort struct {
	h *Repo
}

func TestHomePageHandler(t *testing.T) {
	// tmplcach := forumapp.NewTemplateCache()
	// tmplcach.CreateTemplatesCache()
	// fapp := forumapp.NewForumApp()
	// fapp.Tmpls = tmplcach
	// h := &Repo{app: fapp}
	// h := &Repo{}
	// h.app.Tmpls.CreateTemplatesCache() // Pages["home.page.html"]=
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	fmt.Println(req)
	// h.HomePageHandler(writer, req)
	if writer.Code != http.StatusOK {
		t.Errorf("Expected Method %d,got %d", http.StatusOK, writer.Code)
	}
}

func TestLoginPage(t *testing.T) {
}

func TestSignUpPage(t *testing.T) {
}
