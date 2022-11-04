package response

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func SendErrorData(err *gin.Error, c *gin.Context) {
	c.AbortWithStatusJSON(errorHandler.ConvertError(err.Err), gin.H{"error": err.Error()})
}
