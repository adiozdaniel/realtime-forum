package routes

import (
	"net/http"

	"forum/middlewares"
)

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {

	// ===== Protected RESTFUL API Endpoints ===== //

	// === Posts ===
	mux.Handle("/api/posts/create", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.CreatePost)))
	// mux.Handle("/api/posts/delete", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	// mux.Handle("/api/posts/update", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	// mux.Handle("/api/posts/like", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	// mux.Handle("/api/posts/dislike", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	// === End Posts ===

	// === Comments ===
	// mux.Handle("/api/comments/create", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// mux.Handle("/api/comments/delete", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// mux.Handle("/api/comments/update", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// mux.Handle("/api/comments/like", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// mux.Handle("/api/comments/dislike", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// === End Comments ===

	// === Auth ===
	mux.Handle("/api/auth/check", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.CheckAuth)))
	// === End Auth ===

	// ===== End Protected RESTFUL API Endpoints ===== //

	// ==== Static files server ====
	fs := r.app.Tmpls.GetProjectRoute("/static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fs))))

	// ===== Unprotected RESTFUL API Endpoints ===== //

	// === Posts ===
	mux.HandleFunc("/api/posts", r.postsRepo.AllPosts)
	// === End Posts ===

	// Unprotected Auth RESTFUL API Endpoints
	mux.HandleFunc("/api/auth/register", r.authRepo.RegisterHandler)
	mux.HandleFunc("/api/auth/logout", r.authRepo.LogoutHandler)
	mux.HandleFunc("/api/auth/login", r.authRepo.LoginHandler)
	// ===== End Unprotected RESTFUL API Endpoints =====

	// Page routes
	mux.HandleFunc("/", r.rendersRepo.HomePageHandler)
	mux.HandleFunc("/auth", r.rendersRepo.LoginPageHandler)
	mux.HandleFunc("/auth-sign-up", r.rendersRepo.SignUpPageHandler)

	// CORS middleware
	handler := middlewares.CorsMiddleware(mux)
	return handler
}
