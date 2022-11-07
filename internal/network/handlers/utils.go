package handlers

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func getEmailFromContext(c *gin.Context) (string, error) {
	email, ok := c.Get("userEmail")
	if !ok {
		return "", errorHandler.ErrUnauthorized
	}
	emailStr, ok := email.(string)
	if !ok {
		return "", errorHandler.ErrBadRequest
	}
	return emailStr, nil
}
