package handlers

import (
	"HeadHunter/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := service.GenerateToken(input.Email, input.Password)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func signUp(c *gin.Context) {

}
