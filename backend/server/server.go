package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"forum/forumapp"
	"forum/routes"
)

type Server struct {
	app    *forumapp.ForumApp
	routes *routes.Routes
	server *http.Server
	port   string
}

func NewServer() *Server {
	port := os.Getenv("PORT")
	app, err := forumapp.ForumInit()
	if err != nil {
		log.Fatal(err)
	}

	routes := routes.NewRoutes(app)

	return &Server{
		app, routes, nil, port,
	}
}

func (s *Server) serverSetup() error {
	mux := http.NewServeMux()

	// Register Routes
	h := s.routes.RegisterRoutes(mux)

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: h,
	}

	return nil
}

func (s *Server) Start(ctx context.Context) {
	s.port = s.validatePort()
	s.serverSetup()

	go func() {
		log.Printf("Server starting @http://localhost:%s", s.port)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed:", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server shutting down...")

	s.app.Db.Close()

	// use a timeout context to shut down the server
	clsCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(clsCtx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	} else {
		log.Println("Server successuly shutdown!")
	}
}
