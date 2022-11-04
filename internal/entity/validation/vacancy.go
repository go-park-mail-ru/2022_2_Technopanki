package validation

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

func UpdateVacancyValidate(u models.UpdateVacancy) error {
	if u.JobType == "" && u.Title == "" && u.Description == "" && u.Tasks == "" && u.Requirements == "" && u.Extra == "" && u.Salary == "" && u.Location == "" && u.Experience == "" && u.Format == "" && u.Hours == "" {
		return errorHandler.ErrUpdateStructHasNoValues
	}
	return nil
}

func QueryVacancyValidation(query *gorm.DB) error {
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errorHandler.ErrVacancyNotFound
	}
	return nil
}
