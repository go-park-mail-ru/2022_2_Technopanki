package middleware

import (
	"HeadHunter/internal/network/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Next()
			rawErr := c.Errors.Last()
			response.SendErrorData(rawErr.Err, c)
		}
		return
	}
}
