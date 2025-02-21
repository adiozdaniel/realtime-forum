package routes

import (
	"net/http"
)

// Register routes
func (r *Routes) RegisterRoutes(mux *http.ServeMux) http.Handler {

	// ===== Protected RESTFUL API Endpoints ===== //

	// === Posts ===
	mux.Handle("/api/posts/create", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.CreatePost)))
	// mux.Handle("/api/posts/delete", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	// mux.Handle("/api/posts/update", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.AllPosts)))
	mux.Handle("/api/posts/like", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.PostAddLike)))
	mux.Handle("/api/posts/dislike", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.PostDislike)))
	mux.Handle("/api/posts/image", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.UploadPostImage)))
	// === End Posts ===

	// === Comments ===
	mux.Handle("/api/posts/comments/create", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.CreatePostComment)))
	// mux.Handle("/api/comments/delete", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// mux.Handle("/api/comments/update", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	mux.Handle("/api/comments/like", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.CommentAddLike)))
	// mux.Handle("/api/comments/dislike", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.AllPosts)))
	// === End Comments ===

	// === Replies ===
	mux.Handle("/api/comments/reply/create", r.auth.AuthMiddleware(http.HandlerFunc(r.postsRepo.CreatePostReply)))
	// === End Replies

	// === Auth ===
	mux.Handle("/api/auth/uploadProfilePic", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.UploadProfilePic)))
	mux.Handle("/api/user/dashboard", r.auth.AuthMiddleware(http.HandlerFunc(r.authRepo.UserDashboard)))
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
	mux.HandleFunc("/api/auth/check", r.authRepo.CheckAuth)
	mux.HandleFunc("/api/auth/register", r.authRepo.RegisterHandler)
	mux.HandleFunc("/api/auth/logout", r.authRepo.LogoutHandler)
	mux.HandleFunc("/api/auth/login", r.authRepo.LoginHandler)
	// ===== End Unprotected RESTFUL API Endpoints =====

	// Page routes
	mux.HandleFunc("/", r.rendersRepo.HomePageHandler)
	mux.HandleFunc("/auth", r.rendersRepo.LoginPageHandler)
	mux.HandleFunc("/dashboard", r.rendersRepo.ProfilePageHandler)
	mux.HandleFunc("/auth-sign-up", r.rendersRepo.SignUpPageHandler)

	// CORS middleware
	handler := r.auth.CorsMiddleware(r.auth.UserContextMiddleware(mux))
	return handler
}
