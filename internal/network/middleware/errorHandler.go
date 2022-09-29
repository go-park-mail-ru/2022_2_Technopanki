package middleware

import "github.com/gin-gonic/gin"

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}
