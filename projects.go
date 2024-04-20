package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errProjectNameRequired = errors.New("project name is required")

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

	if err := validateProjectPayload(project); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateProject(project)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating project"})
	}

	WriteJson(w, http.StatusCreated, t)

}

func (s *ProjectsService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t, err := s.store.GetProject(id)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "project not found"})
		return
	}
	WriteJson(w, http.StatusOK, t)
}

func validateProjectPayload(project *Project) error {
	if project.Name == "" {
		return errProjectNameRequired
	}

	return nil
}
