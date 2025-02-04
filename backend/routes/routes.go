package routes

import (
	"net/http"

	"forum/handlers"
	"forum/middlewares"
)

// Register routes
func RegisterRoutes(mux *http.ServeMux) http.Handler {
	mux.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/api/posts", handlers.PostsHandler)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
