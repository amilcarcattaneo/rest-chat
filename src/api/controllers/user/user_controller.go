package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"rest-chat/src/api/controllers"
	"rest-chat/src/api/domain"
	userService "rest-chat/src/api/services/user"
)

// CreateUser handler to register a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}

	id, err := userService.CreateUser(newUser)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}

	newUserIDResponse := domain.NewUserIDResponse{
		ID: id,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newUserIDResponse)
}

// LoginUser handler to login an user
func LoginUser(w http.ResponseWriter, r *http.Request) {

	var userLogin domain.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := userService.LoginUser(userLogin)
	if err != nil {
		controllers.HandleError(w, http.StatusUnauthorized, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse)
}

// AuthenticatedUser handler to check if an user is authenticated
func AuthenticatedUser(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if strings.Contains(tokenString, "Bearer") || strings.Contains(tokenString, "bearer") {
		tokenSplitted := strings.SplitAfter(tokenString, " ")
		if len(tokenSplitted) < 1 {
			controllers.HandleError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
		tokenString = tokenSplitted[1]
	}

	authenticated, err := userService.AuthenticatedUser(tokenString)
	if err != nil {
		controllers.HandleError(w, http.StatusUnauthorized, err)
		return
	}

	if authenticated.Authenticated {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&authenticated)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}
