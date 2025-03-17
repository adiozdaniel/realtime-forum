package forumapp

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type TemplateCache struct {
	Pages map[string]*template.Template
}

func newTemplateCache() *TemplateCache {
	return &TemplateCache{Pages: make(map[string]*template.Template)}
}

func (t *TemplateCache) GetProjectRoute(paths ...string) string {
	cwd, _ := os.Getwd()
	allPaths := append([]string{cwd}, paths...)

	return filepath.Join(allPaths...)
}

func (t *TemplateCache) CreateTemplatesCache() error {
	tmpDir := t.GetProjectRoute("templates", "*.page.html")

	pages, err := filepath.Glob(tmpDir)
	if err != nil {
		return fmt.Errorf("internal server error: adding page")
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return fmt.Errorf("internal server error: parsing page")
		}

		layoutsPath := t.GetProjectRoute("templates", "*.layout.html")

		matches, err := filepath.Glob(layoutsPath)
		if err != nil {
			return fmt.Errorf("internal server error: finding layout page")
		}

		if len(matches) > 0 {
			tpl, err = tpl.ParseGlob(layoutsPath)
			if err != nil {
				return fmt.Errorf("internal server error: parsing layout files")
			}
		}

		t.Pages[name] = tpl
	}
	return nil
}

func (t *TemplateCache) GetPage(name string) (*template.Template, error) {
	tpl, ok := t.Pages[name]
	if !ok {
		return nil, fmt.Errorf("the page '%s' is missing", name)
	}

	return tpl, nil
}
