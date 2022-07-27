package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmrbcnto/go-net-http/entities"
	"github.com/hmrbcnto/go-net-http/middlewares"
)

func (userHandler *user_http_handler) createUser(w http.ResponseWriter, r *http.Request) {
	// Get http request body
	user := new(entities.User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	createdUser, err := userHandler.userUsecase.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter.Encode(createdUser)
}

func (userHandler *user_http_handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Get all users
	w.Header().Set("Content-Type", "application/json")
	users, err := userHandler.userUsecase.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(users)
}

func (userHandler *user_http_handler) getUserById(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	// Get params
	params := mux.Vars(r)
	userId := params["id"]

	user, err := userHandler.userUsecase.GetUserById(userId)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter.Encode(user)
}

func (userHandler *user_http_handler) InitRoutes(mux *mux.Router) {
	subRouter := mux.PathPrefix("").Subrouter()

	// Generate routes
	subRouter.Use(middlewares.CheckForToken)
	subRouter.HandleFunc("/users", userHandler.getAllUsers).Methods("GET")
	subRouter.HandleFunc("/users/id", userHandler.getUserById).Methods("GET")
	subRouter.HandleFunc("/users", userHandler.createUser).Methods("POST")
}
