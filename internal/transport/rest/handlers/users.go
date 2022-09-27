package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"main.go/internal/entities"
	"main.go/internal/service/auth"
	"main.go/internal/storage"
	"net/http"
	"os"
	"strings"
)

func Auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	key, exists := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println("no field ACCESS_TOKEN_SECRET in .env")
		return
	}

	var handler = auth.Handler{}
	user, err := handler.ParseToken(strings.Split(authHeader, " ")[1], []byte(key))
	if err != nil {
		if err.Error() == "invalid access token" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusBadRequest)
	}

	if user, err = storage.FindUserByEmail(user.Email); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func Signup(c *gin.Context) {
	var authHandler auth.Handler
	var requestUser = entities.User{}

	if err := c.BindJSON(&requestUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	storage.CreateUser(requestUser)

	requestUserJWT, err := authHandler.Auth(requestUser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"JWT": requestUserJWT,
	})
}
