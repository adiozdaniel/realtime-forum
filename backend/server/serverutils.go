package server

import (
	"bufio"
	"context"
	"log"
	"os"
)

// ServerCommands listens for user input and handles commands.
func (s *Server) ServerCommands(cancel context.CancelFunc) {
	cmd := bufio.NewScanner(os.Stdin)

	for cmd.Scan() {
		switch cmd.Text() {
		case "exit":
			cancel()
			return
		case "help":
			log.Println("Available commands:\n- 'exit': Shutdown the server\n- 'help': Show this message")
		default:
			log.Println("Unrecognized command. Type 'help' for available commands.")
		}
	}

	if err := cmd.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
		cancel() // Ensure graceful shutdown even on error
	}
}
