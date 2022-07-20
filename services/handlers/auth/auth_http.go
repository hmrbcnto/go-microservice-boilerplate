package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/go-net-http/services/usecases"
)

type AuthHTTPHandler interface {
	InitRoutes(mux *mux.Router)
}

type auth_http_handler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHandler(authUsecase usecases.AuthUsecase) AuthHTTPHandler {
	return &auth_http_handler{
		authUsecase: authUsecase,
	}
}
