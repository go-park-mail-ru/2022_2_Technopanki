package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler            UserH
	VacancyHandler         VacancyH
	VacancyActivityHandler VacancyActivityH
	ResumeHandler          ResumeH
	cfg                    *configs.Config
}

func NewHandlers(usecases *usecases.UseCases, _cfg *configs.Config, _sr session.Repository, handlers *Handlers) *Handlers {
	return &Handlers{
		UserHandler:            newUserHandler(usecases, _cfg, _sr),
		VacancyHandler:         newVacancyHandler(usecases, handlers),
		VacancyActivityHandler: newVacancyActivityHandler(usecases, handlers),
	}
}

type UserH interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	Logout(c *gin.Context)
	AuthCheck(c *gin.Context)
	UpgradeUser(c *gin.Context)
	GetUserId(c *gin.Context) uint
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
}

type ResumeH interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
