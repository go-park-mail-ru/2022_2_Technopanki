package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"fmt"
	"gorm.io/gorm"
)

func notFound(object string) error {
	switch object {
	case "user":
		return errorHandler.ErrUserNotExists
	case "vacancy":
		return errorHandler.ErrVacancyNotFound
	case "vacancy_activity":
		return errorHandler.ErrCannotApplyForVacancy
	case "resume":
		return errorHandler.ErrResumeNotFound
	default:
		return fmt.Errorf("%s not found", object)
	}
}

func QueryValidation(query *gorm.DB, object string) error {
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			return notFound(object)
		}
		return fmt.Errorf("postgre query error: %s", query.Error.Error())
	}
	if query.RowsAffected == 0 {
		return notFound(object)
	}
	return nil
}

func GetUser(email string, db *gorm.DB) (*models.UserAccount, error) {
	var result models.UserAccount
	query := db.Where("email = ?", email).Find(&result)
	return &result, QueryValidation(query, "user")
}
