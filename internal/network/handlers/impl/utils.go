package impl

import (
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
)

func getEmailFromContext(c *gin.Context) (string, error) {
	email, ok := c.Get("userEmail")
	if !ok {
		return "", errorHandler.ErrUnauthorized
	}
	emailStr, ok := email.(string)
	if !ok {
		return "", errorHandler.ErrBadRequest
	}
	return emailStr, nil
}

func (uh *UserHandler) GetUserId(c *gin.Context) (uint, error) {
	email, emailErr := getEmailFromContext(c)
	if emailErr != nil {
		return 0, emailErr
	}
	userId, userIdErr := uh.userUseCase.GetUserId(email)
	if userIdErr != nil {
		_ = c.Error(userIdErr)
		return 0, userIdErr
	}
	if userId == 0 {
		_ = c.Error(errorHandler.ErrUserNotExists)
		return 0, errorHandler.ErrUserNotExists
	}
	return userId, nil
}
