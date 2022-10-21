package handlers

import (
	"HeadHunter/configs"
	"HeadHunter/internal/network/middleware"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler    UserH
	VacancyHandler VacancyH
	ResumeHandler  ResumeH

	SessionMiddleware middleware.SessionMiddleware
	cfg               *configs.Config
}

func NewHandlers(usecases *usecases.UseCases, _cfg *configs.Config, _sr session.Repository) *Handlers {
	return &Handlers{
		UserHandler:       newUserHandler(usecases, _cfg, _sr),
		SessionMiddleware: middleware.NewSessionMiddleware(_sr),
	}
}

type UserH interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	Logout(c *gin.Context)
	AuthCheck(c *gin.Context)
}

type VacancyH interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ResumeH interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
