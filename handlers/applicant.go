package handlers

import (
	jobflow "HeadHunter"
	"HeadHunter/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetApplicants(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.IndentedJSON(http.StatusOK, jobflow.Applicants)
	} else {
		for _, a := range jobflow.Applicants {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
	}
}
func PostApplicants(c *gin.Context) {
	var newApplicant entity.Applicant
	if err := c.BindJSON(&newApplicant); err != nil {
		return
	}
	jobflow.Applicants = append(jobflow.Applicants, newApplicant)
	c.IndentedJSON(http.StatusCreated, newApplicant)
}

func GetApplicantByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range jobflow.Applicants {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Applicant not found"})
}
