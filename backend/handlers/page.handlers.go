package handlers

import (
	"net/http"
)

// HomePage handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	fs := "../frontend/index.html"
	http.ServeFile(w, r, fs)
	return
}
