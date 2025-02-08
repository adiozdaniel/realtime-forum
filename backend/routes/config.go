package routes

import (
	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories/authrepo"
	"forum/repositories/commentrepo"
	"forum/repositories/postrepo"
	"forum/repositories/renders"
)

type Routes struct {
	app          *forumapp.ForumApp
	auth         *middlewares.AuthContext
	authRepo     *authrepo.AuthRepo
	postsRepo    *postrepo.PostsRepo
	rendersRepo  *renders.RendersRepo
	commentsRepo *commentrepo.CommentRepo
}

func NewRoutes(
	app *forumapp.ForumApp,
) *Routes {
	auth := middlewares.NewAuthContext(app)
	authRepo := authrepo.NewAuthRepo(app)
	postsRepo := postrepo.NewPostsRepo(app)
	rendersRepo := renders.NewRendersRepo(app)
	commentsRepo := commentrepo.NewCommentRepo(app)

	return &Routes{
		app,
		auth,
		authRepo,
		postsRepo,
		rendersRepo,
		commentsRepo,
	}
}
