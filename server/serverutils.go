package server

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// ServerCommands listens for user input and handles commands.
func (s *Server) ServerCommands(cancel context.CancelFunc) {
	cmd := bufio.NewScanner(os.Stdin)

	for cmd.Scan() {
		switch strings.TrimSpace(cmd.Text()) {
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

// isPortInUse checks if a given port is already in use.
func (s *Server) isPortInUse(port string) bool {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true // Port is in use
	}
	l.Close() // Release the port
	return false
}

// validatePort checks if the provided port is valid and finds a free port if necessary.
func (s *Server) validatePort() string {
	if s.port == "" {
		s.port = "4000" // Default port
	}

	numPort, err := strconv.Atoi(s.port)
	if err != nil || numPort < 1024 || numPort > 65535 {
		log.Printf("Invalid port '%s', defaulting to 4000", s.port)
		s.port = "4000"
		numPort = 4000
	}

	// If port 4000 is in use, increment until a free one is found
	for {
		if !s.isPortInUse(strconv.Itoa(numPort)) {
			break
		}

		// Port is in use, try the next one
		log.Printf("Port %d is in use. Trying next available port...", numPort)
		numPort++

		// Avoid going out of range
		if numPort > 65535 {
			log.Fatal("No available ports found in the valid range (1024-65535). Exiting.")
		}
	}

	return strconv.Itoa(numPort)
}
