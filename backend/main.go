package main

import (
	"log"
	"net/http"
	"os"

	"forum/handlers"
	"forum/routes"
)

func main() {
	handlers.InitDB() // Initialize SQLite database connection
	mux := http.NewServeMux()

	// Register Routes
	h := routes.RegisterRoutes(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Server starting @http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
