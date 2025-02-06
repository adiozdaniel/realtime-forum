package handlers

import (
	"net/http"

	"forum/forumapp"
)

type Repo struct {
	app *forumapp.ForumApp
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	return &Repo{app}
}

// HomePage handler
func (h *Repo) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := h.app.Tmpls.GetPage("home.page.html")
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{} {
		"Page": "home",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}

// Login page
func (h *Repo) LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := h.app.Tmpls.GetPage("login.page.html")
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}
	// data := map[string]interface{} {
	// 	"isAuthPage": true,
	// }

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}

// sign-up page
func (h *Repo) SignUpPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := h.app.Tmpls.GetPage("signup.page.html")
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{} {
	// 	"isAuthPage": true,
	// }

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}

