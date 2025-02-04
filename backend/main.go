package main

import (
	"log"
	"net/http"
	"os"

	"forum/handlers"
)

func main() {
	handlers.InitDB() // Initialize SQLite database connection

	// Create server
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/api/posts", handlers.PostsHandler)

	// CORS middleware
	handler := corsMiddleware(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting @http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
