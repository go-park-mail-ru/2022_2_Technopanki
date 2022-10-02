package middleware

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Next()
			rawErr := c.Errors.Last()
			c.AbortWithStatusJSON(errorHandler.ConvertError(rawErr.Err), gin.H{"error:": rawErr.Error()})
		}
		return
	}
}
