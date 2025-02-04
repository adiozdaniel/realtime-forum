package forumapp

import (
	"html/template"
	"sync"
)

var (
	instance *ForumApp
	once sync.Once
)

func ForumInit() *ForumApp{
	var err error

	once.Do(func() {
		instance = &ForumApp{
			Tmpls: &TemplateCache {make(map[string]*template.HTML)},
			Errors: &err,
		} 
	})
	return instance
}
