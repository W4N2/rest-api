package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectsService struct {
	store Store
}

func NewProjectsService(s Store) *ProjectsService {
	return &ProjectsService{store: s}
}

func (s *ProjectsService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", s.handleCreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", s.handleGetProject).Methods("GET")
}

func (s *ProjectsService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "invalid Request payload"})
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "invalid Request payload"})
		return
	}

}
