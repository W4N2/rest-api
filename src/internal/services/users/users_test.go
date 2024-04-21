package users

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

func TestNewUserService(t *testing.T) {
	t.Run("should return a new instance of UserService", func(t *testing.T) {
		ms := &store.MockStore{}
		service := NewUserService(ms)

		if service == nil {
			t.Error("expected an instance of UserService, but got nil")
		}
	})
}

func TestCreateUser(t *testing.T) {
	ms := &store.MockStore{}
	service := NewUserService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := types.User{
			Email: "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", service.handleUserRegister)

		router.ServeHTTP(rr, req)

		// Test should return Bad Request.
		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code, it should fail")
		}

	})
}

func TestGetUser(t *testing.T) {
	ms := &store.MockStore{}
	service := NewUserService(ms)

	t.Run("should return the user", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users/{id}", service.handleGetUser)

		router.ServeHTTP(rr, req)

		// Test should return Bad Request.
		if rr.Code != http.StatusOK {
			t.Error("invalid status code, it should fail")
		}
	})
}

func TestCreateAndSetAuthCookie(t *testing.T) {
	ms := &store.MockStore{}
	service := NewUserService(ms)

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := types.User{
			Email: "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", service.handleUserRegister)

		router.ServeHTTP(rr, req)

		// Test should return Bad Request.
		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code, it should fail")
		}

	})
}

func TestValidateUserPayload(t *testing.T) {
	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := types.User{
			Email: "",
		}

		err := validateUserPayload(&payload)
		if err == nil {
			t.Error("expected an error, but got nil")
		}
	})

	t.Run("should return an error if password is empty", func(t *testing.T) {
		payload := types.User{
			Email:    "	",
			Password: "",
		}

		err := validateUserPayload(&payload)
		if err == nil {
			t.Error("expected an error, but got nil")
		}
	})

	t.Run("should return nil if name, email and password are not empty", func(t *testing.T) {
		payload := types.User{
			FirstName: "test",
			LastName:  "test",
			Email:     "test@test.com",
			Password:  "password",
		}

		err := validateUserPayload(&payload)
		if err != nil {
			t.Error("expected nil, but got an error")
		}
	})
}
