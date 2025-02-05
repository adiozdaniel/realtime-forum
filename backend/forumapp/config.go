package forumapp

import (
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
			Tmpls:  &TemplateCache{make(map[string]*template.HTML)},
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
