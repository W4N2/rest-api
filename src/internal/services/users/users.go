package users

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w4n2/rest-api/src/auth"
	"github.com/w4n2/rest-api/src/config"
	"github.com/w4n2/rest-api/src/internal/types"
	"github.com/w4n2/rest-api/src/internal/utils"
	"github.com/w4n2/rest-api/src/store"
)

var errEmailRequired = errors.New("email is required")
var errUsersNameRequired = errors.New("fisrt name and last name are required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/{id}", s.handleGetUser).Methods("GET")

}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	defer r.Body.Close()

	var user *types.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "invalid Request payload"})
		return
	}

	if err := validateUserPayload(user); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "error creating user"})
		return
	}
	user.Password = hashedPassword

	u, err := s.store.CreateUser(user)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, utils.ErrorResponse{Error: "error creating user sesion"})
		return
	}
	utils.WriteJson(w, http.StatusCreated, token)
}

func (s *UserService) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t, err := s.store.GetUserByID(id)
	if err != nil {
		// Return generic response
		utils.WriteJson(w, http.StatusBadRequest, utils.ErrorResponse{Error: "invalid request payload"})
		return
	}
	utils.WriteJson(w, http.StatusOK, t)
}

func validateUserPayload(user *types.User) error {
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
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, id)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
