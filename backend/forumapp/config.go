package forumapp

import (
	"sync"
)

type ForumApp struct {
	Tmpls  *TemplateCache
	Errors error
}

func newForumApp() *ForumApp {
	return &ForumApp{
		Tmpls: newTemplateCache(),
	}
}

var (
	instance *ForumApp
	once     sync.Once
)

func ForumInit() (*ForumApp, error) {
	var err error
	once.Do(func() {
		instance = newForumApp()
	})

	err = instance.Tmpls.CreateTemplatesCache()
	if err != nil {
		return nil, err
	}

	return instance, nil
}
