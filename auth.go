package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

// Higher order func for request auth

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the request
		tokenString := GetTokenFromRequest(r)
		// validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to authenticate token")
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permissioned denied").Error()})
			return
		}

		if !token.Valid {
			log.Printf("failed to authenticate token")
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permissioned denied").Error()})
			return
		}
		// get the userID from the token
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		_, err = store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user")
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permissioned denied").Error()})
			return
		}
		// call the handler func and continue to endpont
		handlerFunc(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func validateJWT(t string) (*jwt.Token, error) {
	secret := Envs.JWTSecret
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

}
