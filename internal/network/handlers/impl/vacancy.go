package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VacancyHandler struct {
	vacancyUseCase usecases.Vacancy
}

func NewVacancyHandler(useCases *usecases.UseCases) *VacancyHandler {
	return &VacancyHandler{vacancyUseCase: useCases.Vacancy}
}

func (vh *VacancyHandler) GetAllVacancies(c *gin.Context) {
	vacancies, getAllErr := vh.vacancyUseCase.GetAll()
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	c.JSON(http.StatusOK, models.GetAllVacanciesResponcePointer{
		Data: vacancies,
	})
}

func (vh *VacancyHandler) GetVacancyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	vacancy, err := vh.vacancyUseCase.GetById(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, vacancy)

}

func (vh *VacancyHandler) GetUserVacancies(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	vacancies, GetErr := vh.vacancyUseCase.GetByUserId(uint(companyId))
	if GetErr != nil {
		_ = c.Error(GetErr)
		return
	}
	c.JSON(http.StatusOK, models.GetAllVacanciesResponcePointer{
		Data: vacancies,
	})

}

func (vh *VacancyHandler) CreateVacancy(c *gin.Context) {
	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.Vacancy
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	id, err := vh.vacancyUseCase.Create(email, &input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (vh *VacancyHandler) DeleteVacancy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	deleteErr := vh.vacancyUseCase.Delete(email, uint(id))
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.Status(http.StatusOK)
}

func (vh *VacancyHandler) UpdateVacancy(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}

	email, contextErr := getEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	var input models.Vacancy
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := vh.vacancyUseCase.Update(email, uint(id), &input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}
	c.Status(http.StatusOK)
}
