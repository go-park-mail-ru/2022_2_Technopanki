package validation

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
)

func UpdateVacancyValidate(u models.UpdateVacancy) error {
	if u.JobType == "" && u.Title == "" && u.Description == "" && u.Tasks == "" && u.Requirements == "" && u.Extra == "" && u.Salary == "" && u.Location == "" && u.Experience == "" && u.Format == "" && u.Hours == "" {
		return errorHandler.ErrUpdateStructHasNoValues
	}
	return nil
}
