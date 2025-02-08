package repositories

import (
	"forum/forumapp"
	"forum/response"
)

type Repo struct {
	app *forumapp.ForumApp
	res *response.JSONRes
}

func NewRepo(app *forumapp.ForumApp) *Repo {
	return &Repo{
		app, response.NewJSONRes(),
	}
}
