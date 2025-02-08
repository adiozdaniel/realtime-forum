package routes

import (
	"net/http"

	"forum/forumapp"
	"forum/middlewares"
	"forum/repositories"
)

type Routes struct {
	app  *forumapp.ForumApp
	repo *repositories.Repo
}

func NewRoutes(app *forumapp.ForumApp, repo *repositories.Repo) *Routes {
	return &Routes{
		app:  app,
		repo: repo,
	}
}

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {
	// Auth routes
	auth := middlewares.NewAuthContext(r.app)
	mux.Handle("/api/posts", auth.AuthMiddleware(http.HandlerFunc(r.repo.PostsHandler)))
	mux.Handle("/api/auth/check", auth.AuthMiddleware(http.HandlerFunc(r.repo.CheckAuth)))

	// Page routes
	fs := r.app.Tmpls.GetProjectRoute("/static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fs))))

	mux.HandleFunc("/api/auth/register", r.repo.RegisterHandler)
	mux.HandleFunc("/api/auth/logout", r.repo.LogoutHandler)
	mux.HandleFunc("/api/auth/login", r.repo.LoginHandler)
	mux.HandleFunc("/auth", r.repo.LoginPage)
	mux.HandleFunc("/auth-sign-up", r.repo.SignUpPage)
	mux.HandleFunc("/", r.repo.HomePageHandler)
	mux.HandleFunc("/api/profile", r.repo.ProfilePage)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
