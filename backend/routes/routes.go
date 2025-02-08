package routes

import (
	"net/http"

	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories/authrepo"
	"forum/repositories/commentrepo"
	"forum/repositories/postrepo"
	"forum/repositories/renders"
)

type Routes struct {
	app          *forumapp.ForumApp
	authRepo     *authrepo.AuthRepo
	postsRepo    *postrepo.PostsRepo
	rendersRepo  *renders.RendersRepo
	commentsRepo *commentrepo.CommentRepo
}

func NewRoutes(
	app *forumapp.ForumApp,
) *Routes {
	authRepo := authrepo.NewAuthRepo(app)
	postsRepo := postrepo.NewPostsRepo(app)
	rendersRepo := renders.NewRendersRepo(app)
	commentsRepo := commentrepo.NewCommentRepo(app)

	return &Routes{
		app,
		authRepo,
		postsRepo,
		rendersRepo,
		commentsRepo,
	}
}

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {
	// Auth routes
	auth := middlewares.NewAuthContext(r.app)
	mux.Handle("/api/posts", auth.AuthMiddleware(http.HandlerFunc(r.authRepo.PostsHandler)))
	mux.Handle("/api/auth/check", auth.AuthMiddleware(http.HandlerFunc(r.authRepo.CheckAuth)))

	// Page routes
	fs := r.app.Tmpls.GetProjectRoute("/static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fs))))

	mux.HandleFunc("/api/auth/register", r.authRepo.RegisterHandler)
	mux.HandleFunc("/api/auth/logout", r.authRepo.LogoutHandler)
	mux.HandleFunc("/api/auth/login", r.authRepo.LoginHandler)
	mux.HandleFunc("/auth", r.rendersRepo.LoginPageHandler)
	mux.HandleFunc("/auth-sign-up", r.rendersRepo.SignUpPageHandler)
	mux.HandleFunc("/", r.rendersRepo.HomePageHandler)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
