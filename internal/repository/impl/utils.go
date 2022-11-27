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

func FilterQueries(filterName string) string {
	switch filterName {
	case "Title":
		return "title LIKE ?"
	case "Location":
		return "location = ?"
	case "Format":
		return "format = ?"
	case "Experience":
		return "experience = ?"
	case "FirstSalaryValue":
		return "salary BETWEEN ? AND ?"
	default:
		return ""
	}
}

func FilterQueryStringFormatter(queryString string, argsSlice []string, db *gorm.DB) *gorm.DB {
	switch len(argsSlice) {
	case 1:
		return db.Where(queryString, argsSlice[0])
	case 2:
		return db.Where(queryString, argsSlice[0], argsSlice[1])
	case 3:
		return db.Where(queryString, argsSlice[0], argsSlice[1], argsSlice[2])
	case 4:
		return db.Where(queryString, argsSlice[0], argsSlice[1], argsSlice[2], argsSlice[3])
	case 5:
		return db.Where(queryString, argsSlice[0], argsSlice[1], argsSlice[2], argsSlice[3], argsSlice[4])
	case 6:
		return db.Where(queryString, argsSlice[0], argsSlice[1], argsSlice[2], argsSlice[3], argsSlice[4], argsSlice[5])
	default:
		return nil
	}
}
