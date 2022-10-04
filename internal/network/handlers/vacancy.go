package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
		c.IndentedJSON(http.StatusOK, jobflow.Vacancies)
	} else {
		if elem, ok := jobflow.Vacancies[id]; ok {
			c.IndentedJSON(http.StatusOK, elem)
			return
		} else {
			_ = c.Error(errorHandler.ErrVacancyNotFound)
		}
	}
}
