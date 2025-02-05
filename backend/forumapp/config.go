package forumapp

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sync"
)

var (
	instance *ForumApp
	once     sync.Once
)

func ForumInit() *ForumApp {
	var err error

	once.Do(func() {
		instance = &ForumApp{
			Tmpls:  &TemplateCache{make(map[string]*template.Template)},
			Errors: &err,
		}
	})
	return instance
}

func (app *ForumApp) getProjectRoute(paths ...string) string {
	cwd, _ := os.Getwd()
	allPaths := append([]string{cwd}, paths...)

	return filepath.Join(allPaths...)
}

func (app *ForumApp) createTemplatesCache() error {
	tmpDir := app.getProjectRoute("frontend/templates", "*.page.html")

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

		layoutsPath := app.getProjectRoute("frontend/templates", "*.layout.html")

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

		app.Tmpls.Pages[name] = tpl
	}

	return nil
}
