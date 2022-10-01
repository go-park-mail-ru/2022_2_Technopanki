package middleware

import (
	"HeadHunter/internal/network/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Добавить mutex, userSession убрать, шифрование пароля пользователя
func Session(c *gin.Context) {
	sessionToken, err := c.Cookie("session")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userSession, err := sessions.SessionsStore.GetSession(sessions.Token(sessionToken))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if userSession.IsExpired() {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
