package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrBadRequest        = errors.New("bad request")
	ErrEmptyHeader       = errors.New("empty header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrUnauthorized      = errors.New("unauthorized")
)

var errorToCode = map[error]int{
	ErrBadRequest:        http.StatusBadRequest,
	ErrEmptyHeader:       http.StatusUnauthorized,
	ErrInvalidAuthHeader: http.StatusUnauthorized,
	http.ErrNoCookie:     http.StatusUnauthorized,
	ErrUnauthorized:      http.StatusUnauthorized,
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
