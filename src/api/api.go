package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w4n2/rest-api/src/internal/services/projects"
	"github.com/w4n2/rest-api/src/internal/services/tasks"
	"github.com/w4n2/rest-api/src/internal/services/users"
	"github.com/w4n2/rest-api/src/store"
)

type APIServer struct {
	addr  string
	store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering services
	TasksService := tasks.NewTasksService(s.store)
	TasksService.RegisterRoutes(subrouter)

	ProjectsService := projects.NewProjectsService(s.store)
	ProjectsService.RegisterRoutes(subrouter)

	UserService := users.NewUserService(s.store)
	UserService.RegisterRoutes(subrouter)

	log.Println("Starting API Server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
