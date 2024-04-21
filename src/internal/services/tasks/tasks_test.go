package tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/w4n2/rest-api/src/internal/types"
	"github.com/w4n2/rest-api/src/store"
)

func TestCreateTask(t *testing.T) {
	ms := &store.MockStore{}
	service := NewTasksService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := types.Task{
			Name: "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask)

		router.ServeHTTP(rr, req)

		// Test should return Bad Request.
		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code, it should fail")
		}

	})
}

func TestGetTask(t *testing.T) {
	ms := &store.MockStore{}
	service := NewTasksService(ms)

	t.Run("should return the task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks/{id}", service.handleGetTask)

		router.ServeHTTP(rr, req)

		// Test should return Bad Request.
		if rr.Code != http.StatusOK {
			t.Error("invalid status code, it should fail")
		}
	})
}
