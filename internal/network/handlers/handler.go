package handlers

import (
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler    UserH
	VacancyHandler VacancyH
	ResumeHandler  ResumeH
}

func NewHandlers(usecases *usecases.UseCases) *Handlers {
	return &Handlers{
		UserHandler: newUserHandler(usecases),
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
