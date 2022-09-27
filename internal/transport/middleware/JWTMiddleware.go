package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"main.go/internal/service/auth"
	"net/http"
	"os"
	"strings"
)

func JWTMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var handler = auth.Handler{}
	key, exists := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println("no field ACCESS_TOKEN_SECRET in .env")
		return
	}

	_, err := handler.ParseToken(headerParts[1], []byte(key))
	if err != nil {
		if err.Error() == "invalid access token" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusBadRequest)
	}
}
