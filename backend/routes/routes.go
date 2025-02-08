package routes

import (
	"net/http"

	"forum/middlewares"
)

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {

	mux.Handle("/api/posts", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.PostsHandler)))
	mux.Handle("/api/auth/check", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.CheckAuth)))

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
