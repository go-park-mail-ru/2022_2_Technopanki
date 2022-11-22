package errorHandler

import (
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Next()
			rawErr := c.Errors.Last()
			Response(c, rawErr.Err)
		}
		return
	}
}
