package service

import (
	"HeadHunter/database"
	"HeadHunter/errorHandler"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTTL   = 24 * time.Hour
	signingKey = "flasvuw,qdpaskgnqwoasmflqwkfmaoq"
)

type tokenClaims struct {
	jwt.StandardClaims
	ID string
}

func GenerateToken(email, password string) (string, error) {
	employer, err := database.GetEmployer(email, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		employer.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return errorHandler.ReturnErrorCase[interface{}]("invalid signing method")
			}

			return []byte(signingKey), nil
		})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return errorHandler.ReturnErrorCase[string]("token claims are not of type *tokenClaims")
	}
	return claims.ID, nil
}
