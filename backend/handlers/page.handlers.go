package handlers

import (
	"fmt"
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

	// Fetch the template
	tmpl, err := t.GetPage("home.page.html")
	if err != nil {
		// Log the error for debugging
		fmt.Println(err)
		// Send the internal server error response
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	// Execute the template only if no error occurred
	err = tmpl.Execute(w, nil)
	if err != nil {
		// If template execution fails, log the error and return a 500 internal server error
		fmt.Println(err)
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}
