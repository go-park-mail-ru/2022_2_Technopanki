package impl

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/network/response"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	cfg         *configs.Config
	userUseCase usecases.User
}

func NewUserHandler(useCases *usecases.UseCases, _cfg *configs.Config) *UserHandler {
	return &UserHandler{cfg: _cfg, userUseCase: useCases.User}
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
	if !input.TwoFactorSignIn {
		c.SetCookie("session", token, uh.cfg.DefaultExpiringSession, "/", uh.cfg.Domain,
			uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
		response.SendSuccessData(c, &input)
	} else {
		c.Status(http.StatusAccepted)
		return
	}
}

func (uh *UserHandler) SignUp(c *gin.Context) {
	var input models.UserAccount
	if err := easyjson.UnmarshalFromReader(c.Request.Body, &input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	token, signUpErr := uh.userUseCase.SignUp(&input)
	if signUpErr != nil {
		_ = c.Error(signUpErr)
		return
	}
	if !uh.cfg.Security.ConfirmAccountMode {
		c.SetCookie("session", token, uh.cfg.DefaultExpiringSession, "/", uh.cfg.Domain,
			uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
		c.SetSameSite(http.SameSiteLaxMode)
		response.SendSuccessData(c, &input)
	} else {
		c.Status(http.StatusOK)
	}

}

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
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}
	user, err := uh.userUseCase.AuthCheck(email)
	if err != nil {
		_ = c.Error(err)
		return
	}
	response.SendSuccessData(c, user)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}
	var input models.UserAccount
	if err := easyjson.UnmarshalFromReader(c.Request.Body, &input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	input.Email = email

	updateErr := uh.userUseCase.UpdateUser(&input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}
	response.SendSuccessData(c, &input)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
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
	userJson, errJson := user.MarshalJSON()
	if errJson != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (uh *UserHandler) GetAllEmployers(c *gin.Context) {
	var employers []*models.UserAccount
	var getAllErr error
	var filters models.UserFilter

	search := c.Query("search")
	if search != "" {
		filters.CompanyName = search
	}

	workField := c.Query("field")
	if workField != "" {
		filters.BusinessType = workField
	}

	size := c.Query("size")
	if size != "" {
		if strings.Contains(size, ":") {
			split := strings.Split(size, ":")
			filters.FirstCompanySizeValue = split[0]
			filters.SecondCompanySizeValue = split[1]
		} else {
			_ = c.Error(errorHandler.ErrBadRequest)
			return
		}
	}

	city := c.Query("city")
	if city != "" {
		filters.Location = city
	}

	employers, getAllErr = uh.userUseCase.GetAllEmployers(filters)
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	var usersResponse models.GetAllUsersResponcePointer
	usersResponse.Data = employers
	employersJson, err := usersResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", employersJson)
}

func (uh *UserHandler) GetAllApplicants(c *gin.Context) {
	var applicants []*models.UserAccount
	var getAllErr error
	var filters models.UserFilter

	search := c.Query("search")
	if search != "" {
		if strings.Index(search, " ") != -1 {
			split := strings.Split(search, " ")
			filters.ApplicantName = split[0]
			filters.ApplicantSurname = split[1]
		} else {
			filters.ApplicantSurname = search
		}
	}

	city := c.Query("city")
	if city != "" {
		filters.Location = city
	}

	age := c.Query("age")
	if age != "" {
		if strings.Contains(age, ":") {
			split := strings.Split(age, ":")
			filters.FirstAgeValue = split[0]
			filters.SecondAgeValue = split[1]
		} else {
			_ = c.Error(errorHandler.ErrBadRequest)
			return
		}
	}

	applicants, getAllErr = uh.userUseCase.GetAllApplicants(filters)
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	var usersResponse models.GetAllUsersResponcePointer
	usersResponse.Data = applicants
	applicantsJson, err := usersResponse.MarshalJSON()
	if err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", applicantsJson)
}

func (uh *UserHandler) GetUserSafety(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	user, getErr := uh.userUseCase.GetUserSafety(uint(id))
	if getErr != nil {
		_ = c.Error(getErr)
		return
	}
	userJson, errJson := user.MarshalJSON()
	if errJson != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (uh *UserHandler) UploadUserImage(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	user, getUserErr := uh.userUseCase.GetUserByEmail(email)

	if getUserErr != nil {
		_ = c.Error(errorHandler.ErrUserNotExists)
		return
	}

	_, file, fileErr := c.Request.FormFile("avatar")
	if fileErr != nil {
		_ = c.Error(errorHandler.ErrInvalidFileFormat)
		return
	}
	_, uploadErr := uh.userUseCase.UploadUserImage(user, file)
	if uploadErr != nil {
		_ = c.Error(uploadErr)
		return
	}
	response.SendUploadImageData(c, user)
}

func (uh *UserHandler) DeleteUserImage(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
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

func (uh *UserHandler) GetPreview(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	user, getUserErr := uh.userUseCase.GetUserSafety(uint(id))
	if getUserErr != nil {
		_ = c.Error(getUserErr)
		return
	}
	response.SendPreviewData(c, user)
}

func (uh *UserHandler) ConfirmUser(c *gin.Context) {
	var input struct {
		Code  string `json:"code"`
		Email string `json:"email"`
	}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	user, token, confirmErr := uh.userUseCase.ConfirmUser(input.Code, input.Email)
	if confirmErr != nil {
		_ = c.Error(confirmErr)
		return
	}

	c.SetCookie("session", token, uh.cfg.DefaultExpiringSession, "/",
		uh.cfg.Domain, uh.cfg.Cookie.Secure, uh.cfg.Cookie.HTTPOnly)
	response.SendSuccessData(c, user)
}

func (uh *UserHandler) UpdatePassword(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	var input struct {
		Code     string `json:"code"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updatePasswordErr := uh.userUseCase.UpdatePassword(input.Code, email, input.Password)

	if updatePasswordErr != nil {
		_ = c.Error(updatePasswordErr)
		return
	}

	c.Status(http.StatusOK)
}

func (uh *UserHandler) GetMailing(c *gin.Context) {
	email, emailErr := utils.GetEmailFromContext(c)
	if emailErr != nil {
		_ = c.Error(emailErr)
		return
	}

	err := uh.userUseCase.GetMailing(email)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
