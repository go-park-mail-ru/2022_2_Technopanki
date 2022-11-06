package response

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func SendErrorData(c *gin.Context, err error) {
	result := gin.H{"error": err.Error(), "descriptors": errorHandler.GetErrorDescriptors(err)}
	c.AbortWithStatusJSON(errorHandler.ConvertError(err), result)
}
