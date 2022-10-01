package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrBadRequest         = errors.New("bad request")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrUserExists         = errors.New("user already exists")
)

var errorToCode = map[error]int{
	ErrBadRequest:         http.StatusBadRequest,
	ErrUnauthorized:       http.StatusUnauthorized,
	ErrServiceUnavailable: http.StatusServiceUnavailable,
	ErrUserExists:         http.StatusBadRequest,
}

func ConvertError(err error) int {
	result, ok := errorToCode[err]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Next()
			rawErr := c.Errors.Last()
			c.AbortWithStatusJSON(ConvertError(rawErr.Err), gin.H{"error:": rawErr.Error()})
		}
		return
	}
}
