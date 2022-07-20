package usecases

import (
	"github.com/hmrbcnto/go-net-http/entities"
	"github.com/hmrbcnto/go-net-http/infastructure/db/mongo/auth_repo"
)

type AuthUsecase interface {
	Login(username string, password string) (*entities.User, error)
}

type authUsecase struct {
	authRepo auth_repo.AuthRepo
}

func NewAuthUsecase(authRepo auth_repo.AuthRepo) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (authUsecase *authUsecase) Login(username string, password string) (*entities.User, error) {
	// Hash password input here

	user, err := authUsecase.authRepo.Login(username, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
