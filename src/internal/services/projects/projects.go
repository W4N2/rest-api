package projects

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w4n2/rest-api/src/auth"
	"github.com/w4n2/rest-api/src/internal/types"
	"github.com/w4n2/rest-api/src/internal/utils"
	"github.com/w4n2/rest-api/src/store"
)

var errProjectNameRequired = errors.New("project name is required")

type ProjectsService struct {
	store store.Store
}

func NewProjectsService(s store.Store) *ProjectsService {
	return &ProjectsService{store: s}
}

func (s *ProjectsService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", auth.WithJWTAuth(s.handleCreateProject, s.store)).Methods("POST")
	r.HandleFunc("/projects/{id}", auth.WithJWTAuth(s.handleGetProject, s.store)).Methods("GET")

}

func (s *ProjectsService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	defer r.Body.Close()

	var project *types.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	if err := validateProjectPayload(project); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateProject(project)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "error creating project"})
	}

	utils.WriteJson(w, http.StatusCreated, t)

}

func (s *ProjectsService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t, err := s.store.GetProject(id)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: "project not found"})
		return
	}
	utils.WriteJson(w, http.StatusOK, t)
}

func validateProjectPayload(project *types.Project) error {
	if project.Name == "" {
		return errProjectNameRequired
	}

	return nil
}
