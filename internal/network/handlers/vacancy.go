package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

type VacancyHandler struct {
	vacancyUseCase usecases.Vacancy
	userHandler    UserH
}

func newVacancyHandler(useCases *usecases.UseCases, handlers *Handlers) *VacancyHandler {
	return &VacancyHandler{vacancyUseCase: useCases.Vacancy, userHandler: handlers.UserHandler}
}

type getAllVacanciesResponce struct {
	Data []models.Vacancy `json:"data"`
}

func (vh *VacancyHandler) GetAll(c *gin.Context) {
	vacancies, getAllErr := vh.vacancyUseCase.GetAll()
	if getAllErr != nil {
		_ = c.Error(getAllErr)
		return
	}
	c.JSON(http.StatusOK, getAllVacanciesResponce{
		vacancies,
	})
}

func (vh *VacancyHandler) GetById(c *gin.Context) {
	userId := vh.userHandler.GetUserId(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidId)
		return
	}
	vacancy, err := vh.vacancyUseCase.GetById(userId, id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, vacancy)

}

func (vh *VacancyHandler) Create(c *gin.Context) {
	userId := vh.userHandler.GetUserId(c)
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

func (vh *VacancyHandler) Delete(c *gin.Context) {
	userId := vh.userHandler.GetUserId(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidId)
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

func (vh *VacancyHandler) Update(c *gin.Context) {
	userId := vh.userHandler.GetUserId(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidId)
		return
	}
	var input models.UpdateVacancy
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

// @Summary GetVacancies
// @Tags Получить вакансии
// @Description Получить вакансии
// @ID get-vacancies
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Vacancy
// @Failure 400 {body} string "invalid query"
// @Failure 404 {body} string "vacancy not found"
// @Router /api/vacancy/ [get]
func GetVacancies(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil && idStr != "" {
		_ = c.Error(errorHandler.ErrInvalidQuery)
		return
	}
	if id == 0 {
		outputSlice := make([]entity.Vacancy, 0, len(jobflow.Vacancies))
		keys := make([]int, 0)
		for k, _ := range jobflow.Vacancies {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			outputSlice = append(outputSlice, jobflow.Vacancies[k])
		}
		c.IndentedJSON(http.StatusOK, outputSlice)
	} else {
		if elem, ok := jobflow.Vacancies[id]; ok {
			c.IndentedJSON(http.StatusOK, elem)
			return
		} else {
			_ = c.Error(errorHandler.ErrVacancyNotFound)
		}
	}
}
