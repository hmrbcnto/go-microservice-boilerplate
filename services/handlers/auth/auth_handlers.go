package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmrbcnto/go-net-http/entities"
)

func (authHandler *auth_http_handler) login(w http.ResponseWriter, r *http.Request) {
	// Get request body
	loginData := new(entities.LoginStruct)
	err := json.NewDecoder(r.Body).Decode(loginData)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set header
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)

	loginResults, err := authHandler.authUsecase.Login(loginData.Username, loginData.Password)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	jsonWriter.Encode(loginResults)
}

func (authHandler *auth_http_handler) InitRoutes(mux *mux.Router) {
	// Generate routes
	mux.HandleFunc("/login", authHandler.login).Methods("POST")
}
