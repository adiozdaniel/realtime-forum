package renders

import (
	"net/http"
)

// HomePage handler
func (m *RendersRepo) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	data := map[string]interface{}{
		"Page": "home",
	}

	err := m.RenderTemplate(w, "home.page.html", data)
	if err != nil {
		http.Error(w, "couldnt get the template", http.StatusInternalServerError)
	}
}

// Login page
func (m *RendersRepo) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	err := m.RenderTemplate(w, "login.page.html", nil)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}

// sign-up page
func (m *RendersRepo) SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	err := m.RenderTemplate(w, "signup.page.html", nil)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
		return
	}
}

// moderator page
func (m *RendersRepo) ModeratorPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	data := map[string]interface{}{
		"Page": "moderator",
	}

	err := m.RenderTemplate(w, "moderator.page.html", data)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
		return
	}
}

// admin page
func (m *RendersRepo) AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	data := map[string]interface{}{
		"Page": "admin",
	}

	err := m.RenderTemplate(w, "admin.page.html", data)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
		return
	}
}

// ProfilePageHandler renders profile page
func (m *RendersRepo) ProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	data := map[string]interface{}{
		"Page": "profile",
	}

	err := m.RenderTemplate(w, "profile.page.html", data)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
		return
	}
}

func (m *RendersRepo) NotFoundPageHandler( w http.ResponseWriter, r *http.Request) {
	err := m.RenderTemplate(w, "pageNotFound.page.html", nil )
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
		return
	}
}
