package renders

import (
	"net/http"
)

// RenderTemplate renders an HTML template
func (m *RendersRepo) RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	tmpl, err := m.app.Tmpls.GetPage(templateName)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}
