package routes

import (
	"net/http"

	"forum/forumapp"
	"forum/handlers"
	"forum/middlewares"
)

type Routes struct {
	app  *forumapp.ForumApp
	repo *handlers.Repo
}

func NewRoutes(app *forumapp.ForumApp, repo *handlers.Repo) *Routes {
	return &Routes{
		app:  app,
		repo: repo,
	}
}

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {
	// Page routes
	fs := r.app.Tmpls.GetProjectRoute("/static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fs))))

	mux.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/api/posts", handlers.PostsHandler)
	mux.HandleFunc("/", r.repo.HomePageHandler)
	mux.HandleFunc("/auth", r.repo.LoginPage)
	mux.HandleFunc("/auth-sign-up", r.repo.SignUpPage)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
