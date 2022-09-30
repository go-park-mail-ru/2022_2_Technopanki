package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	var input = signInInput{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

}

//func setNewTokenInCookie(c *gin.Context) string {
//	token := uuid.NewString()
//	c.SetCookie("session", token, 100, "/", "localhost", false, false)
//
//	return token
//}

var store = sessions{}

func SignUp(c *gin.Context) {
	var input = signUpInput{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token := SessionsMap.NewSession(input.Email)

	c.SetCookie("session", token, int(time.Hour*10), "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}
