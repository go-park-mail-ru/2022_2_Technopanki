package handlers

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, err := h.uc.User.SignIn(&input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", viper.GetString("domain"), false, true)
	c.JSON(http.StatusOK, gin.H{"name": input.Name, "surname": input.Surname})
}

func (h *Handler) SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, signUpErr := h.uc.User.SignUp(input)
	if signUpErr != nil {
		_ = c.Error(signUpErr)
		return
	}
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", viper.GetString("domain"), false, true)
	c.Status(http.StatusOK)
}

// @Summary      Logout
// @Description  Выход пользователя
// @Tags         Авторизация
// @ID logout
// @Accept       json
// @Produce      json
// @Success 200
// @Failure 400 {body} string "bad request"
// @Router       /auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	logoutErr := h.uc.User.Logout(token)
	if logoutErr != nil {
		_ = c.Error(logoutErr)
		return
	}
	c.SetCookie("session", token, -1, "/", viper.GetString("domain"), false, true)
}

func (h *Handler) AuthCheck(c *gin.Context) {
	email, ok := c.Get("userEmail")
	if !ok {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	user, err := h.uc.User.AuthCheck(email.(string))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, entity.User{Name: user.Name, Surname: user.Surname})
}
