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

	tmpl, err := m.app.Tmpls.GetPage("home.page.html")
	if err != nil {
		http.Error(w, "Oops, something went wrong!", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Page": "home",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Oops, something went wrong while rendering the page!", http.StatusInternalServerError)
	}
}

// Login page
func (m *RendersRepo) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := m.app.Tmpls.GetPage("login.page.html")
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
func (m *RendersRepo) SignUpPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Oops, didn't understand what you are looking for", http.StatusForbidden)
		return
	}

	tmpl, err := m.app.Tmpls.GetPage("signup.page.html")
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
