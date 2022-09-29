package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		errorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	//ID, err := jwt.ParseToken(headerParts[1])
	//if err != nil {
	//	errorResponse(c, http.StatusUnauthorized, err.Error())
	//}
	//c.Set("userID", ID)
}
