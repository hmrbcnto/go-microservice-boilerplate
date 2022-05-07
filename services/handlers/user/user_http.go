package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hmrbcnto/go-net-http/services/usecases"
)

type UserHTTPHandler interface {
	InitRoutes(mux *mux.Router)
}

type user_http_handler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) UserHTTPHandler {
	return &user_http_handler{
		userUsecase: userUsecase,
	}
}
