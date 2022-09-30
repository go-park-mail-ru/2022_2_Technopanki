package handlers

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var input = signInInput{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if storage.UserStorage.IsUserInStorage(input.Email) {
		token := sessions.SessionsStore.NewSession(input.Email)
		c.SetCookie("session", token, 100, "/", "localhost", false, false)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	input.ID = uuid.NewString()
	storage.UserStorage.AddUser(input)

	token := sessions.SessionsStore.NewSession(input.Email)

	expiration := time.Now().Add(time.Hour * 24 * 3)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   true,
	}
	http.SetCookie(c.Writer, &cookie)

	fmt.Println(c.Request)
	//c.SetCookie("newCookie", token, 3600, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}
