package entities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginReturn struct {
	User User
	TokenString string
	ExpirationTime time.Time
}