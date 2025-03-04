package forumapp

import (
	"sync"
)

type ForumApp struct {
	Tmpls         *TemplateCache
	Db            *DataConfig
	dbPath        string
	AllowedRoutes map[string]bool
	Errors        error
}

func newForumApp() *ForumApp {
	return &ForumApp{
		Tmpls:         newTemplateCache(),
		Db:            NewDb(),
		dbPath:        "./forum.db",
		AllowedRoutes: make(map[string]bool),
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

		err = instance.Db.InitDB(instance.dbPath)
		err = instance.Tmpls.CreateTemplatesCache()
	})

	instance.AllowedRoutes = map[string]bool{
		"/api/posts/create":          true,
		"/api/posts/delete":          true,
		"/api/posts/like":            true,
		"/api/posts/dislike":         true,
		"/api/posts/image":           true,
		"/api/posts/comments/create": true,
		"/api/comments/delete":       true,
		"/api/comments/update":       true,
		"/api/comments/like":         true,
		"/api/comments/dislike":      true,
		"/api/comments/reply/create": true,
		"/api/auth/uploadProfilePic": true,
		"/api/user/dashboard":        true,
		"/api/user/editBio":          true,
		"/api/notifications/check":   true,
		"/static/":                   true,
		"/api/posts":                 true,
		"/api/auth/check":            true,
		"/api/auth/register":         true,
		"/api/auth/logout":           true,
		"/api/auth/login":            true,
		"/":                          true,
		"/auth":                      true,
		"/dashboard":                 true,
		"/auth-sign-up":              true,
		"/moderator":                 true,
		"/admin":                     true,
	}

	if err != nil {
		return nil, err
	}

	return instance, nil
}
