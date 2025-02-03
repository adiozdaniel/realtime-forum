package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting @http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
