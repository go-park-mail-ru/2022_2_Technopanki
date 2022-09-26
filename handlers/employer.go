package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmployers(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusOK, jobflow.Employers)
	} else {
		for _, a := range jobflow.Employers {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
	}
}
func PostEmployers(c *gin.Context) {
	var newEmployer entity.Employer
	if err := c.BindJSON(&newEmployer); err != nil {
		return
	}
	jobflow.Employers = append(jobflow.Employers, newEmployer)
	c.IndentedJSON(http.StatusCreated, newEmployer)
	return
}

func GetEmployerByID(c *gin.Context) {
	id := c.Param("id")

	for _, e := range jobflow.Employers {
		if e.ID == id {
			c.IndentedJSON(http.StatusOK, e)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employer not found"})
}
