package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	_, check := c.Get("userEmail")
	if !check {
		return
	}
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil && idStr != "" {
		_ = c.Error(errorHandler.ErrInvalidQuery)
		return
	}
	if id == 0 {
		outputSlice := make([]entity.Vacancy, 0, len(jobflow.Vacancies))
		for _, elem := range jobflow.Vacancies {
			outputSlice = append(outputSlice, elem)
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
