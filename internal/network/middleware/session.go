package middleware

import (
	"HeadHunter/internal/repository/session"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct {
	sessionRepos session.Repository
}

func NewSessionMiddleware(sr session.Repository) *SessionMiddleware {
	return &SessionMiddleware{sessionRepos: sr}
}

func (sm *SessionMiddleware) Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}

	userEmail, err := sm.sessionRepos.GetSession(sessionToken)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Set("userEmail", userEmail)
}
