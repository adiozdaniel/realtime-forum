package main

import (
	"context"
	"os"

	"forum/server"
)

func main() {
	port := os.Getenv("PORT")

	s := server.NewServer(port)
	ctx, cancel := context.WithCancel(context.Background())

	go s.ServerCommands(cancel)
	s.Start(ctx)
}
