package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("Name is required")
var errProjectIDRequired = errors.New("Project ID is required")
var errUserIDRequired = errors.New("User ID is required")

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Invalid Request Payload"})
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Invalid Request Payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJson(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.assignedTo == 0 {
		return errUserIDRequired
	}

	return nil
}
