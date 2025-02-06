package server

import (
	"log"
	"net/http"

	"forum/forumapp"
	"forum/handlers"
	"forum/routes"
)

type Server struct {
	app  *forumapp.ForumApp
	repo *handlers.Repo
	routes *routes.Routes
	port string
}

func NewServer(port string) *Server {
	app, _ := forumapp.ForumInit()
	repo := handlers.NewRepo(app)
	routes := routes.NewRoutes(app, repo)
	return &Server{
		app:  app,
		repo: repo,
		routes: routes,
		port: port,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	// Register Routes
	h := s.routes.RegisterRoutes(mux)
	_, err := forumapp.ForumInit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server starting @http://localhost:%s", s.port)
	err = http.ListenAndServe(":"+s.port, h)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
