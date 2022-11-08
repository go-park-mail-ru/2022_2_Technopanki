package handlers

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

func (uh *UserHandler) GetUserType(c *gin.Context) (string, error) {
	email, emailErr := getEmailFromContext(c)
	if emailErr != nil {
		return "", emailErr
	}
	user, userIdErr := uh.userUseCase.GetUserByEmail(email)
	if userIdErr != nil {
		_ = c.Error(userIdErr)
		return "", userIdErr
	}
	return user.UserType, nil
}

func (rh *ResumeHandler) isResumeAvailable(c *gin.Context, id uint) error {
	userId, userErr := rh.userHandler.GetUserId(c)
	if userErr != nil {
		return userErr
	}

	resume, getResumeErr := rh.resumeUseCase.GetResume(id)
	if getResumeErr != nil {
		return getResumeErr
	}

	if resume.UserAccountId != userId {
		return errorHandler.ErrForbidden
	}
	return nil
}
