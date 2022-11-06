package response

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func SendErrorData(err error, c *gin.Context) {
	result := gin.H{"error": err.Error()}
	
	complexErr, ok := err.(*errorHandler.ComplexError)
	if ok {
		desc := complexErr.GetDescriptors()
		result["descriptors"] = desc
	}

	c.AbortWithStatusJSON(errorHandler.ConvertError(err), result)
}
