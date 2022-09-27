package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"main.go/internal/entities"
	"os"
	"strconv"
	"time"
)

type Service interface {
	Auth(user entities.User) (string, error)
	ParseToken(accessToken string, authKey []byte) (string, error)
}

type Handler struct {
	service Service
}

func getTokenTTL() (time.Duration, error) {
	tokenTTLInHoursStr, exists := os.LookupEnv("TOKEN_TTL_IN_HOURS")
	if !exists {
		return 0, errors.New("no field TOKEN_TTL_IN_HOURS in .env")
	}

	tokenTTLInHours, err := strconv.Atoi(tokenTTLInHoursStr)
	if err != nil {
		return 0, err
	}

	return time.Duration(tokenTTLInHours), nil
}

func (a *Handler) Auth(user entities.User) (string, error) {
	key, exists := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !exists {
		return "", errors.New("no field ACCESS_TOKEN_SECRET in .env")
	}

	tokenTTLInHours, err := getTokenTTL()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * tokenTTLInHours)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		User: user,
	})

	return token.SignedString([]byte(key))
}

func (a *Handler) ParseToken(accessToken string, authKey []byte) (entities.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected token")
		}
		return authKey, nil
	})

	if err != nil {
		return entities.User{}, err
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		return claims.User, nil
	}

	return entities.User{}, errors.New("invalid access token")
}
