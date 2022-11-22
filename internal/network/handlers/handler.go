package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/network/handlers/impl"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler            UserH
	VacancyHandler         VacancyH
	VacancyActivityHandler VacancyActivityH
	ResumeHandler          ResumeH
}

func NewHandlers(usecases *usecases.UseCases, _cfg *configs.Config) *Handlers {
	return &Handlers{
		UserHandler:            impl.NewUserHandler(usecases, _cfg),
		ResumeHandler:          impl.NewResumeHandler(usecases),
		VacancyHandler:         impl.NewVacancyHandler(usecases),
		VacancyActivityHandler: impl.NewVacancyActivityHandler(usecases),
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
}

type VacancyH interface {
	GetAllVacancies(c *gin.Context)
	GetVacancyById(c *gin.Context)
	GetUserVacancies(c *gin.Context)
	CreateVacancy(c *gin.Context)
	UpdateVacancy(c *gin.Context)
	DeleteVacancy(c *gin.Context)
}

type VacancyActivityH interface {
	ApplyForVacancy(c *gin.Context)
	GetAllVacancyApplies(c *gin.Context)
	GetAllUserApplies(c *gin.Context)
	DeleteUserApply(c *gin.Context)
}

type ResumeH interface {
	GetResume(c *gin.Context)
	GetResumeByApplicant(c *gin.Context)
	GetPreviewResumeByApplicant(c *gin.Context)
	CreateResume(c *gin.Context)
	UpdateResume(c *gin.Context)
	DeleteResume(c *gin.Context)
}
