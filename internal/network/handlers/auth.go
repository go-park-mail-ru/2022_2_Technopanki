package handlers

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/network/middleware"
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
		_ = c.Error(middleware.ErrBadRequest)
		return
	}

	user := storage.UserStorage.FindByEmail(input.Email)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		_ = c.Error(middleware.ErrUnauthorized)
		return
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", "localhost", false, true)
}

func SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(middleware.ErrBadRequest)
		return
	}

	input.ID = uuid.NewString()
	user := storage.UserStorage.FindByEmail(input.Email)
	if user.Email == input.Email {
		_ = c.Error(middleware.ErrUserExists)
		return
	}

	err := storage.UserStorage.AddUser(input)
	if err != nil {
		_ = c.Error(middleware.ErrServiceUnavailable)
		return
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(middleware.ErrBadRequest)
		return
	}
	c.SetCookie("session", token, -1, "/", "localhost", false, true)
}
