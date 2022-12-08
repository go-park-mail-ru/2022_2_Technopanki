package impl

import (
	"HeadHunter/pkg/errorHandler"
	"errors"
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
	case "vacancy_applies":
		return errorHandler.ErrVacancyApplyNotFound
	default:
		return fmt.Errorf("%s not found", object)
	}
}

func QueryValidation(query *gorm.DB, object string) error {
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return notFound(object)
		}
		return fmt.Errorf("postgre query error: %s", query.Error.Error())
	}
	if query.RowsAffected == 0 {
		return notFound(object)
	}
	return nil
}

func FilterQueryStringFormatter(queryString string, argsSlice []interface{}, db *gorm.DB) *gorm.DB {
	return db.Where(queryString, argsSlice...)
}
