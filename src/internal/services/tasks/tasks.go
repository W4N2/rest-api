package tasks

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

var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("project ID is required")
var errUserIDRequired = errors.New("user ID is required")

type TasksService struct {
	store store.Store
}

func NewTasksService(s store.Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", auth.WithJWTAuth(s.handleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", auth.WithJWTAuth(s.handleGetTask, s.store)).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	defer r.Body.Close()

	var task *types.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "error creating task"})
		return
	}

	utils.WriteJson(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t, err := s.store.GetTask(id)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: "task not found"})
		return
	}

	utils.WriteJson(w, http.StatusOK, t)
}

func validateTaskPayload(task *types.Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedTo == 0 {
		return errUserIDRequired
	}

	return nil
}
