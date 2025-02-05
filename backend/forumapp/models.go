package forumapp

import (
	"html/template"
)

type ForumApp struct {
	Tmpls *TemplateCache
	Errors *error
}

type TemplateCache struct {
	Pages map[string]*template.HTML
}
