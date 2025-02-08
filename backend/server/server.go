package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"forum/forumapp"
	"forum/repositories"
	"forum/routes"
)

type Server struct {
	app    *forumapp.ForumApp
	repo   *repositories.Repo
	routes *routes.Routes
	server *http.Server
	port   string
}

func NewServer(port string) *Server {
	app, err := forumapp.ForumInit()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewRepo(app)
	routes := routes.NewRoutes(app, repo)
	return &Server{
		app:    app,
		repo:   repo,
		routes: routes,
		port:   port,
	}
}

func (s *Server) runServer() error {
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
	s.runServer()

	go func() {
		log.Printf("Server starting @http://localhost:%s", s.port)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed:", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server shutting down...")

	// use a timeout context to shut down the server
	clsCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(clsCtx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	} else {
		log.Println("Server successuly shutdown!")
	}
}
