package response

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func SendErrorData(err *errorHandler.ComplexError, c *gin.Context) {
	result := gin.H{"error": err.Error()}
	desc, descErr := err.GetDescriptors("type")
	if descErr == nil {
		result["type"] = desc
	}
	c.AbortWithStatusJSON(errorHandler.ConvertError(err), result)
}
