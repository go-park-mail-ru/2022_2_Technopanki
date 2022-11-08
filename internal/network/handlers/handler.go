package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler    UserH
	VacancyHandler VacancyH
	ResumeHandler  ResumeH
	cfg            *configs.Config
}

func NewHandlers(usecases *usecases.UseCases, _cfg *configs.Config, _sr session.Repository) *Handlers {
	userHandler := newUserHandler(usecases, _cfg, _sr)
	return &Handlers{
		UserHandler:   userHandler,
		ResumeHandler: newResumeHandler(usecases, _cfg, userHandler),
	}
}

type UserH interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	Logout(c *gin.Context)
	AuthCheck(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserSafety(c *gin.Context)
	UploadUserImage(c *gin.Context)
	DeleteUserImage(c *gin.Context)
	GetPreview(c *gin.Context)
	GetUserId(c *gin.Context) (uint, error)
	GetUserType(c *gin.Context) (string, error)
}

type VacancyH interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ResumeH interface {
	GetResume(c *gin.Context)
	GetResumeByApplicant(c *gin.Context)
	CreateResume(c *gin.Context)
	UpdateResume(c *gin.Context)
	DeleteResume(c *gin.Context)
}
