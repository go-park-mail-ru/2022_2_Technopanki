package middleware

import (
	"HeadHunter/internal/network/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	userSession := sessions.SessionsStore.GetSession(sessions.Token(sessionToken))
	if userSession.IsExpired() {
		newToken := sessions.SessionsStore.UpdateSession(sessions.Token(sessionToken))

		newUserSession := sessions.SessionsStore.GetSession(sessions.Token(newToken))
		if newUserSession.IsExpired() {
			fmt.Println("LOGICAL ERROR IN SESSION MIDDLEWARE")
		}
		c.SetCookie("session", newToken, int(newUserSession.ExpiresAt), "/", "localhost", false, false)
	}
}
