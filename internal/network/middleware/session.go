package middleware

import (
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"github.com/gin-gonic/gin"
)

func Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}

	userSession, err := sessions.SessionsStore.GetSession(sessions.Token(sessionToken))
	if err != nil {
		_ = c.Error(err)
		return
	}
	if userSession.IsExpired() {
		deleteSessionErr := sessions.SessionsStore.DeleteSession(sessions.Token(sessionToken))
		if deleteSessionErr != nil {
			_ = c.Error(deleteSessionErr)
			return
		}
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	c.Set("userEmail", userSession.Email)
}
