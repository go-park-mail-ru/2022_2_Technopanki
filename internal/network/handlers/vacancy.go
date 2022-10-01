package handlers

import (
	jobflow "HeadHunter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVacancies(c *gin.Context) {
	_, check := c.Get("cookie")
	if !check {
		return
	}
	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusOK, jobflow.Vacancies)
	} else {
		for _, v := range jobflow.Vacancies {
			if v.ID == id {
				c.IndentedJSON(http.StatusOK, v)
				return
			}
		}
	}
}

//func PostVacancies(c *gin.Context) {
//	var newVacancy entity.Vacancy
//	if err := c.BindJSON(&newVacancy); err != nil {
//		return
//	}
//	jobflow.Vacancies = append(jobflow.Vacancies, newVacancy)
//	c.IndentedJSON(http.StatusCreated, newVacancy)
//	return
//}
