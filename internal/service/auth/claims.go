package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"main.go/internal/entities"
)

type userClaims struct {
	jwt.RegisteredClaims
	entities.User
}

type VacancyClaims struct {
	jwt.RegisteredClaims
	entities.Vacancy
}
