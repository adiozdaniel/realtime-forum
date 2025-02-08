package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"forum/server"
)

func serverCommands(cancel context.CancelFunc) {
	cmd := bufio.NewScanner(os.Stdin)

	for cmd.Scan() {
		switch cmd.Text() {
		case "exit":
			cancel()
		case "help":
			log.Println("To shutdown the server. Type 'exit")
		default:
			log.Println("Unrecognized command. Type 'help' to see documentation")
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	s := server.NewServer(port)
	ctx, cancel := context.WithCancel(context.Background())

	go serverCommands(cancel)
	s.Start(ctx)
}
