package renders

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders an HTML template
func (m *RendersRepo) RenderTemplate(
	w http.ResponseWriter,
	templateName string,
	data interface{},
) error {
	tmpl, err := m.app.Tmpls.GetPage(templateName)
	if err != nil {
		m.RenderServerError(w)
		return err
	}

	return tmpl.Execute(w, data)
}

func (m *RendersRepo) RenderServerError(w http.ResponseWriter) {
	tmp, err := m.app.Tmpls.GetPage("internalServerError.page.html")
	if err != nil {
		m.RenderTemplate(w, serverError, nil)
		tm, err := template.ParseFiles(serverError)
		if err != nil {
			log.Println(err)
		}

		tm.Execute(w, nil)

		return
	}
	tmp.Execute(w, nil)
}
