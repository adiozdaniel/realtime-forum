package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
}

func TestLoginHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/api/auth/login", nil)
	w := httptest.NewRecorder()
	LoginHandler(w, r)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestPostsHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/api/auth/posts", nil)
	w := httptest.NewRecorder()
	LoginHandler(w, r)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %d got %d", http.StatusMethodNotAllowed, w.Code)
	}
}
