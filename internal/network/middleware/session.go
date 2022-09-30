package middleware

import (
	"HeadHunter/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// token -> user email

type UserSession struct {
	User      entity.User
	ExpiresAt time.Time
}

func (u *UserSession) isExpired() bool {
	return u.ExpiresAt.Unix() <= time.Now().Unix()
}

func (u *UserSession) updateSession() {
	u.ExpiresAt = time.Now().Add(time.Minute)
}

var sessionMap map[string]UserSession

func setNewTokenInCookie(c *gin.Context) string {
	token := uuid.NewString()
	c.SetCookie("session", token, 100, "/", "localhost", false, false)

	return token
}

func Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userSession := sessionMap[sessionToken]
	if userSession.isExpired() {
		delete(sessionMap, sessionToken)

		token := setNewTokenInCookie(c)

		sessionMap[token] = userSession
		userSession.updateSession()
	}
}
