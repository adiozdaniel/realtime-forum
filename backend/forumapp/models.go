package forumapp

import (
	"html/template"
)

type ForumApp struct {
	Tmpls  *TemplateCache
	Errors error
}

type TemplateCache struct {
	Pages map[string]*template.Template
}

func newTemplateCache() *TemplateCache {
	return &TemplateCache{make(map[string]*template.Template)}
}

func newForumApp() *ForumApp{
	return &ForumApp{
		Tmpls: newTemplateCache(),
		Errors: nil,
	}
}
