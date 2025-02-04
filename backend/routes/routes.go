package routes

import (
	"net/http"

	"forum/handlers"
	"forum/middlewares"
)

// Register routes
func RegisterRoutes(mux *http.ServeMux) http.Handler {
	//Page routes
	fs := "frontend"
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(fs))))
	
	mux.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/api/posts", handlers.PostsHandler)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
