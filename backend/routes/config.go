package routes

import (
	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories/authrepo"
	"forum/repositories/postrepo"
	"forum/repositories/renders"
)

type Routes struct {
	app         *forumapp.ForumApp
	auth        *middlewares.AuthContext
	authRepo    *authrepo.AuthRepo
	postsRepo   *postrepo.PostsRepo
	rendersRepo *renders.RendersRepo
}

func NewRoutes(
	app *forumapp.ForumApp,
) *Routes {
	auth := middlewares.NewAuthContext()
	postsRepo := postrepo.NewPostsRepo(app)
	rendersRepo := renders.NewRendersRepo(app)
	authRepo := authrepo.NewAuthRepo(app, auth)

	return &Routes{
		app,
		auth,
		authRepo,
		postsRepo,
		rendersRepo,
	}
}
