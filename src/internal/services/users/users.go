package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errEmailRequired = errors.New("email is required")
var errUsersNameRequired = errors.New("fisrt name and last name are required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/{id}", s.handleGetUser).Methods("GET")

}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "invalid Request payload"})
		return
	}

	defer r.Body.Close()

	var user *User
	err = json.Unmarshal(body, &user)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "invalid Request payload"})
		return
	}

	if err := validateUserPayload(user); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword := HashPassword(user.Password)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating user"})
		return
	}
	user.Password = hashedPassword

	u, err := s.store.CreateUser(user)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating user sesion"})
		return
	}
	WriteJson(w, http.StatusCreated, token)
}

func (s *UserService) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t, err := s.store.GetUserByID(id)
	if err != nil {
		// Return generic response
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}
	WriteJson(w, http.StatusOK, t)
}

func validateUserPayload(user *User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" || user.LastName == "" {
		return errUsersNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
