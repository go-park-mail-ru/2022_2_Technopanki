package handlers

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/sessions"
	"HeadHunter/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// @Summary      SignIn
// @Description  Вход пользователя
// @Tags         Авторизация
// @ID login
// @Accept       json
// @Produce      json
// @Param input body entity.User{} true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Router       /sign-in [post]
func SignIn(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	inputValidity := validation.IsValidateAuthData(input)
	if inputValidity != nil {
		_ = c.Error(inputValidity)
		return
	}
	user, err := storage.UserStorage.FindByEmail(input.Email)
	if err != nil {
		_ = c.Error(err)
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

// @Summary      SignUp
// @Description  Регистрация пользователя
// @Tags         Регистрация
// @ID create-account
// @Accept       json
// @Produce      json
// @Param input body entity.User{} true "account info"
// @Success 200 {string} string "token"
// @Failure 400 {string} string "bad request"
// @Failure 503 {string} string "service unavailable"
// @Router       /sign-up [post]
func SignUp(c *gin.Context) {
	var input = entity.User{}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	inputValidity := validation.IsValidate(input)
	if inputValidity != nil {
		_ = c.Error(inputValidity)
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

// @Summary      Logout
// @Description  Выход пользователя
// @Tags         Авторизация
// @ID logout
// @Accept       json
// @Produce      json
// @Success 200 {string} string "unauthorized"
// @Failure 400 {string} string "bad request"
// @Router       /logout [post]
func Logout(c *gin.Context) {
	token, err := c.Cookie("session")
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	err = sessions.SessionsStore.DeleteSession(sessions.Token(token))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.SetCookie("session", token, -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "unauthorized",
	})
}
