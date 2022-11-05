package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	cfg         *configs.Config
	userUseCase usecases.User
	sessionRepo session.Repository
}

func newUserHandler(useCases *usecases.UseCases, _cfg *configs.Config, _sr session.Repository) *UserHandler {
	return &UserHandler{cfg: _cfg, userUseCase: useCases.User, sessionRepo: _sr}
}
func (uh *UserHandler) SignIn(c *gin.Context) {
	var input models.UserAccount
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, err := uh.userUseCase.SignIn(&input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie("session", token, uh.cfg.DefaultExpiringSession, "/", uh.cfg.Domain,
		uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
	if input.UserType == "applicant" {
		c.JSON(http.StatusOK, gin.H{"name": input.ApplicantName, "surname": input.ApplicantSurname})
	} else if input.UserType == "employer" {
		c.JSON(http.StatusOK, gin.H{"name": input.CompanyName})
	}
}

func (uh *UserHandler) SignUp(c *gin.Context) {
	var input models.UserAccount
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	token, signUpErr := uh.userUseCase.SignUp(&input)
	if signUpErr != nil {
		_ = c.Error(signUpErr)
		return
	}
	c.SetCookie("session", token, uh.cfg.DefaultExpiringSession, "/",
		uh.cfg.Domain, uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
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

	logoutErr := uh.userUseCase.Logout(token)
	if logoutErr != nil {
		_ = c.Error(logoutErr)
		return
	}
	c.SetCookie("session", token, -1, "/",
		uh.cfg.Domain, uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
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
	user, err := uh.userUseCase.AuthCheck(emailStr)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if user.UserType == "applicant" {
		c.JSON(http.StatusOK, gin.H{"name": user.ApplicantName, "surname": user.ApplicantSurname})
	} else if user.UserType == "employer" {
		c.JSON(http.StatusOK, gin.H{"name": user.CompanyName})
	}
}

func (uh *UserHandler) UpgradeUser(c *gin.Context) {
	var input models.UserAccount
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	upgradeErr := uh.userUseCase.UpgradeUser(&input)
	if upgradeErr != nil {
		_ = c.Error(upgradeErr)
		return
	}
	c.Status(http.StatusOK)
}

func (uh *UserHandler) GetUserId(c *gin.Context) (uint, error) {
	userEmail, ok := c.Get("userEmail")
	if !ok {
		paramErr := c.Error(errorHandler.ErrInvalidParam)
		return 0, paramErr
	}
	emailString := userEmail.(string)
	userId, userIdErr := uh.userUseCase.GetUserId(emailString)
	if userIdErr != nil {
		err := c.Error(userIdErr)
		return 0, err
	}
	return userId, nil
}
