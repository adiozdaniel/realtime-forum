package main

import (
	"log"
	"net/http"
	"os"

	"forum/handlers"
	"forum/middlewares"
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
	handler := middlewares.CorsMiddleware(mux)

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
