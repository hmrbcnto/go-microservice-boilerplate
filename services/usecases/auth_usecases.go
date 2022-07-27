package usecases

import (
	// JWT Library

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hmrbcnto/go-net-http/entities"
	"github.com/hmrbcnto/go-net-http/infastructure/db/mongo/auth_repo"
)

type AuthUsecase interface {
	Login(username string, password string) (*entities.LoginReturn, error)
}

type authUsecase struct {
	authRepo auth_repo.AuthRepo
}

func NewAuthUsecase(authRepo auth_repo.AuthRepo) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (authUsecase *authUsecase) Login(username string, password string) (*entities.LoginReturn, error) {
	// Hash password input here

	user, err := authUsecase.authRepo.Login(username, password)

	if err != nil {
		return nil, err
	}

	// Creating JWT Key
	var jwtKey = []byte("secret_key")

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &entities.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, err
	}

	return &entities.LoginReturn{
		User: *user,
		TokenString: tokenString,
		ExpirationTime: expirationTime,
	}, nil
}
