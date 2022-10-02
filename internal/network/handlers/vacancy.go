package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/internal/errorHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVacancies(c *gin.Context) {
	_, check := c.Get("userID")
	if !check {
		return
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		_ = c.Error(errorHandler.ErrInvalidQuery)
	}
	if id == 0 {
		c.IndentedJSON(http.StatusOK, jobflow.Vacancies)
	} else {
		if elem, ok := jobflow.Vacancies[id]; ok {
			c.IndentedJSON(http.StatusOK, elem)
			return
		}
	}
}

//func PostVacancy(c *gin.Context) {
//	var newVacancy entity.Vacancy
//	if err := c.BindJSON(&newVacancy); err != nil {
//		return
//	}
//	jobflow.Vacancies = append(jobflow.Vacancies, newVacancy)
//	c.IndentedJSON(http.StatusCreated, newVacancy)
//	return
//}
