package server

import (
	"context"
	"log"
	"net/http"

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
		<-ctx.Done()
		log.Println("Server shutting down...")
		if err := s.server.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down server: %v\n", err)
		}

		log.Println("Server successuly shutdown!")
		return
	}()

	log.Printf("Server starting @http://localhost:%s", s.port)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Server failed:", err)
	}
}
