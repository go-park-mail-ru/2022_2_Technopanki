package validation

import (
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

func QueryVacancyValidation(query *gorm.DB) error {
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errorHandler.ErrVacancyNotFound
	}
	return nil
}
