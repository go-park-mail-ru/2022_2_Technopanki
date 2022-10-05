package handlers

import (
	jobflow "HeadHunter"
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
// @Param input body entity.User{} true "user data"
// @Success 200 {body} string "OK"
// @Failure 400 {body} string "bad request"
// @Failure 401 {body} string "unauthorized"
// @Router /auth/sign-in [post]
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
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", jobflow.Domain, false, true)
	c.JSON(http.StatusOK, gin.H{"name": user.Name, "surname": user.Surname})
}

// @Summary      SignUp
// @Description  Регистрация пользователя
// @Tags         Регистрация
// @ID create-account
// @Accept       json
// @Produce      json
// @Param input body entity.User{} true "account info"
// @Success 200 {body} string "OK"
// @Failure 400 {body} string "bad request"
// @Failure 503 {body} string "service unavailable"
// @Router  /auth/sign-up [post]
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
	c.SetCookie("session", token, int(sessions.SessionsStore.DefaultExpiresAt), "/", jobflow.Domain, false, true)
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
	c.SetCookie("session", token, -1, "/", jobflow.Domain, false, true)
}

func AuthCheck(c *gin.Context) {
	email, ok := c.Get("userEmail")
	if !ok {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	user, err := storage.UserStorage.FindByEmail(email.(string))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, entity.User{Name: user.Name, Surname: user.Surname})
}
