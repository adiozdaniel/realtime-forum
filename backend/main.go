package main

import (
	"context"

	"forum/server"
)

func main() {
	s := server.NewServer()
	ctx, cancel := context.WithCancel(context.Background())

	go s.ServerCommands(cancel)
	s.Start(ctx)
}
