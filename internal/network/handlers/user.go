package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	cfg         *configs.Config
	userUseCase usecases.User
	sessionRepo session.Repository
}

func newUserHandler(useCases *usecases.UseCases, _cfg *configs.Config, _sr session.Repository) *UserHandler {
	return &UserHandler{cfg: _cfg, userUseCase: useCases.User, sessionRepo: _sr}
}

func (uh *UserHandler) getEmailFromContext(c *gin.Context) (string, error) {
	email, ok := c.Get("userEmail")
	if !ok {
		return "", errorHandler.ErrUnauthorized
	}
	emailStr, ok := email.(string)
	if !ok {
		return "", errorHandler.ErrBadRequest
	}
	return emailStr, nil
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
	email, emailErr := uh.getEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}
	user, err := uh.userUseCase.AuthCheck(email)
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

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	email, emailErr := uh.getEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}
	var input models.UserAccount
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	if email != input.Email {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	updateErr := uh.userUseCase.UpdateUser(&input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}
	c.Status(http.StatusOK)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	email, emailErr := uh.getEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	idStr := c.Query("id")
	id, queryErr := strconv.Atoi(idStr)
	if queryErr != nil {
		_ = c.Error(errorHandler.ErrInvalidQuery)
		return
	}
	user, getErr := uh.userUseCase.GetUser(uint(id))
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	if user.Email != email {
		_ = c.Error(errorHandler.ErrUnauthorized)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUserSafety(c *gin.Context) {
	idStr := c.Query("id")
	id, queryErr := strconv.ParseUint(idStr, 10, 32)
	if queryErr != nil {
		_ = c.Error(errorHandler.ErrInvalidQuery)
		return
	}
	user, getErr := uh.userUseCase.GetUserSafety(uint(id))
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UploadUserImage(c *gin.Context) {
	email, emailErr := uh.getEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	user, getUserErr := uh.userUseCase.GetUserByEmail(email)

	if getUserErr != nil {
		_ = c.Error(errorHandler.ErrUserNotExists)
		return
	}
	form, formErr := c.MultipartForm()
	if formErr != nil {
		_ = c.Error(formErr)
		return
	}
	var fileName string
	var imgExt string
	for key := range form.File {
		fileName = key
		arr := strings.Split(fileName, ".")
		if len(arr) < 2 {
			_ = c.Error(errorHandler.ErrBadRequest)
			return
		}
		imgExt = arr[len(arr)-1]
	}
	file, _, fileErr := c.Request.FormFile(fileName)
	if fileErr != nil {
		_ = c.Error(fileErr)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			_ = c.Error(err)
		}
	}(file)

	uploadErr := uh.userUseCase.UploadUserImage(user, &file, imgExt)
	if uploadErr != nil {
		_ = c.Error(uploadErr)
		return
	}
	c.Status(http.StatusOK)
}

func (uh *UserHandler) DeleteUserImage(c *gin.Context) {
	email, emailErr := uh.getEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	user, getUserErr := uh.userUseCase.GetUserByEmail(email)

	if getUserErr != nil {
		_ = c.Error(errorHandler.ErrUserNotExists)
		return
	}

	deleteErr := uh.userUseCase.DeleteUserImage(user)
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.Status(http.StatusOK)
}
