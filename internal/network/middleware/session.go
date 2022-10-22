package middleware

import (
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository/session"
	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct {
	sr session.Repository
}

func NewSessionMiddleware(sr session.Repository) SessionMiddleware {
	return SessionMiddleware{sr: sr}
}
func (sm *SessionMiddleware) Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}

	userSession, err := sm.sr.GetSession(session.Token(sessionToken))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Set("userEmail", userSession)
}
