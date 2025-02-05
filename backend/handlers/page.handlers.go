package handlers

import (
	"net/http"

	"forum/forumapp"
)

var t forumapp.TemplateCache

// HomePage handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := t.GetPage("home.page.html")
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}
