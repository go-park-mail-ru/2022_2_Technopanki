package utils

import (
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
)

func GetEmailFromContext(c *gin.Context) (string, error) {
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
