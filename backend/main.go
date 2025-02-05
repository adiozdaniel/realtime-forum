package main

import (
	"os"

	"forum/server"
)

func main() {
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	s := server.NewServer(port)
	s.Start()
}
