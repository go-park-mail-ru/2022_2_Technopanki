package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	cfg  *configs.Config
	User usecases.User
}

func newUserHandler(usecases *usecases.UseCases, _cfg *configs.Config) *UserHandler {
	return &UserHandler{cfg: _cfg, User: usecases.User}
}
func (uh *UserHandler) SignIn(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, err := uh.User.SignIn(&input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", uh.cfg.Domain, false, true)
	c.JSON(http.StatusOK, gin.H{"name": input.Name, "surname": input.Surname})
}

func (uh *UserHandler) SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, signUpErr := uh.User.SignUp(input)
	if signUpErr != nil {
		_ = c.Error(signUpErr)
		return
	}
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", uh.cfg.Domain, false, true)
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
func (uh *UserHandler) Logout(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	logoutErr := uh.User.Logout(token)
	if logoutErr != nil {
		_ = c.Error(logoutErr)
		return
	}
	c.SetCookie("session", token, -1, "/", uh.cfg.Domain, false, true)
}

func (uh *UserHandler) AuthCheck(c *gin.Context) {
	email, ok := c.Get("userEmail")
	if !ok {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	emailStr, ok := email.(string)
	if !ok {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	user, err := uh.User.AuthCheck(emailStr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, entity.User{Name: user.Name, Surname: user.Surname})
}
