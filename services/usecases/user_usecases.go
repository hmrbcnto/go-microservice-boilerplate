package usecases

import (
	"github.com/hmrbcnto/go-net-http/entities"
	"github.com/hmrbcnto/go-net-http/infastructure/db/mongo/user_repo"
)

type UserUsecase interface {
	GetAllUsers() ([]entities.User, error)
	GetUserById(userId string) (*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
}

type userUsecase struct {
	userRepo user_repo.UserRepo
}

func NewUserUsecase(userRepo user_repo.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (userUsecase *userUsecase) GetAllUsers() ([]entities.User, error) {
	// Business logic if needed here

	users, err := userUsecase.userRepo.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userUsecase *userUsecase) GetUserById(userId string) (*entities.User, error) {
	// More business logic here if needed

	user, err := userUsecase.userRepo.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userUsecase *userUsecase) CreateUser(user *entities.User) (*entities.User, error) {
	// Business logic again

	user, err := userUsecase.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
