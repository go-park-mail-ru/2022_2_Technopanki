package middleware

import (
	"HeadHunter/internal/network/sessions"
	"github.com/gin-gonic/gin"
)

func Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(ErrUnauthorized)
		return
	}

	userSession, err := sessions.SessionsStore.GetSession(sessions.Token(sessionToken))
	if err != nil {
		_ = c.Error(ErrUnauthorized)
		return
	}
	if userSession.IsExpired() {
		_ = c.Error(ErrUnauthorized)
		return
	}
}
