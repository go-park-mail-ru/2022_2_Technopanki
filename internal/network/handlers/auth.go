package handlers

import (
	"HeadHunter/internal/entity"
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := storage.UserStorage.FindByEmail(input.Email)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := sessions.SessionsStore.NewSession(input.Email)
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", "localhost", false, true)
}

func SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	input.ID = uuid.NewString()
	user := storage.UserStorage.FindByEmail(input.Email)
	if user.Email == input.Email {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	err := storage.UserStorage.AddUser(input)
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.SetCookie("session", token, -1, "/", "localhost", false, true)
}
