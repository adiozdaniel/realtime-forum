package routes

import (
	"net/http"

	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories"
	"forum/repositories/authrepo"
)

type Routes struct {
	app      *forumapp.ForumApp
	repo     *repositories.Repo
	authRepo *authrepo.AuthRepo
}

func NewRoutes(app *forumapp.ForumApp, authRepo *authrepo.AuthRepo, repo *repositories.Repo) *Routes {
	return &Routes{
		app:      app,
		repo:     repo,
		authRepo: authRepo,
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
	mux.HandleFunc("/auth", r.repo.LoginPage)
	mux.HandleFunc("/auth-sign-up", r.repo.SignUpPage)
	mux.HandleFunc("/", r.repo.HomePageHandler)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
