package renders

import (
	"html/template"
	"net/http"
)

// HomePage handler
func (m *RendersRepo) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := template.New("homepage").ParseFiles(
		m.app.Tmpls.GetProjectRoute("templates", "index.html"),
		m.app.Tmpls.GetProjectRoute("templates", "header.html"),
		m.app.Tmpls.GetProjectRoute("templates", "sidebar.html"),
		m.app.Tmpls.GetProjectRoute("templates", "footer.html"),
	)
	if err != nil {
		m.RenderServerError(w)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		m.RenderServerError(w)
	}
}

// Login page
func (m *RendersRepo) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	_ = m.RenderTemplate(w, "login.page.html", nil)
}

// sign-up page
func (m *RendersRepo) SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	_ = m.RenderTemplate(w, "signup.page.html", nil)
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

	_ = m.RenderTemplate(w, "moderator.page.html", data)
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

	_ = m.RenderTemplate(w, "admin.page.html", data)
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

	_ = m.RenderTemplate(w, "profile.page.html", data)
}

func (m *RendersRepo) NotFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.RenderTemplate(w, "pageNotFound.page.html", nil)
}

func (m *RendersRepo) InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.RenderTemplate(w, "internalServerError.page.html", nil)
}
