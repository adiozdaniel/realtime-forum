package forumapp

import (
	"sync"
)

type ForumApp struct {
	Tmpls  *TemplateCache
	Db     *DataConfig
	dbPath string
	Errors error
}

func newForumApp() *ForumApp {
	return &ForumApp{
		Tmpls:  newTemplateCache(),
		Db:     NewDb(),
		dbPath: "./forum.db",
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

	err = instance.Db.InitDB(instance.dbPath)
	if err != nil {
		return nil, err
	}

	err = instance.Tmpls.CreateTemplatesCache()
	if err != nil {
		return nil, err
	}

	return instance, nil
}
