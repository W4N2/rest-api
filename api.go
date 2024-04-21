package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering services
	TasksService := NewTasksService(s.store)
	TasksService.RegisterRoutes(subrouter)

	ProjectsService := NewProjectsService(s.store)
	ProjectsService.RegisterRoutes(subrouter)

	UserService := NewUserService(s.store)
	UserService.RegisterRoutes(subrouter)

	log.Println("Starting API Server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
