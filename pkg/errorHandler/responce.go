package errorHandler

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, err error) {
	result := gin.H{"error": err.Error(), "descriptors": GetErrorDescriptors(err)}
	c.AbortWithStatusJSON(ConvertError(err), result)
}
