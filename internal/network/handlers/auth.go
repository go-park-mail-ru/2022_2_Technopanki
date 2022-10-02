package handlers

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	user, err := storage.UserStorage.FindByEmail(input.Email)
	if err != nil {
		if err.Error() == "user not exists" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", "localhost", false, true)
	c.Status(http.StatusOK)
}

func SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	input.ID = uuid.NewString()
	_, err := storage.UserStorage.FindByEmail(input.Email)
	if err == nil {
		_ = c.Error(errorHandler.ErrUserExists)
		return
	}

	err = storage.UserStorage.AddUser(input)
	if err != nil {
		_ = c.Error(errorHandler.ErrServiceUnavailable)
		return
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", "localhost", false, true)
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	// TODO: error
	err = sessions.SessionsStore.DeleteSession(sessions.Token(token))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.SetCookie("session", token, -1, "/", "localhost", false, true)
	c.AbortWithStatus(http.StatusOK)
}
