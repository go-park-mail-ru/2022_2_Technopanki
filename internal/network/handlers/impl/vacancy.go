package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/network/handlers"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VacancyHandler struct {
	vacancyUseCase usecases.Vacancy
	userHandler    handlers.UserH
}

func NewVacancyHandler(useCases *usecases.UseCases, userHandler *UserHandler) *VacancyHandler {
	return &VacancyHandler{vacancyUseCase: useCases.Vacancy, userHandler: userHandler}
}

type getAllVacanciesResponcePointer struct {
	Data []*models.Vacancy `json:"data"`
}
type getAllVacanciesResponce struct {
	Data []models.Vacancy `json:"data"`
}

func (vh *VacancyHandler) GetAllVacancies(c *gin.Context) {
	vacancies, getAllErr := vh.vacancyUseCase.GetAll()
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	c.JSON(http.StatusOK, getAllVacanciesResponcePointer{
		vacancies,
	})
}

func (vh *VacancyHandler) GetVacancyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	vacancy, err := vh.vacancyUseCase.GetById(id)
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
	vacancies, GetErr := vh.vacancyUseCase.GetByUserId(companyId)
	if GetErr != nil {
		_ = c.Error(GetErr)
		return
	}
	c.JSON(http.StatusOK, getAllVacanciesResponcePointer{
		vacancies,
	})

}

func (vh *VacancyHandler) CreateVacancy(c *gin.Context) {
	userId, getUserIdErr := vh.userHandler.GetUserId(c)
	if getUserIdErr != nil {
		_ = c.Error(getUserIdErr)
		return
	}
	var input models.Vacancy
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	id, err := vh.vacancyUseCase.Create(userId, &input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type statusResponse struct {
	Status string `json:"status"`
}

func (vh *VacancyHandler) DeleteVacancy(c *gin.Context) {
	userId, getUserIdErr := vh.userHandler.GetUserId(c)
	if getUserIdErr != nil {
		_ = c.Error(getUserIdErr)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	deleteErr := vh.vacancyUseCase.Delete(userId, id)
	if deleteErr != nil {
		_ = c.Error(deleteErr)
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (vh *VacancyHandler) UpdateVacancy(c *gin.Context) {
	userId, getUserIdErr := vh.userHandler.GetUserId(c)
	if getUserIdErr != nil {
		_ = c.Error(getUserIdErr)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidParam)
		return
	}
	var input models.Vacancy
	if err := c.BindJSON(&input); err != nil {
		_ = c.Error(errorHandler.ErrBadRequest)
		return
	}

	updateErr := vh.vacancyUseCase.Update(userId, id, &input)
	if updateErr != nil {
		_ = c.Error(updateErr)
		return
	}
	c.Status(http.StatusOK)
}
