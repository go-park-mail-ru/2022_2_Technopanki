package sanitize

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"github.com/microcosm-cc/bluemonday"
	"reflect"
)

var policy = bluemonday.UGCPolicy()

func SanitizeString(data string) (string, error) {
	if data == "" {
		return "", nil
	}
	newString := policy.Sanitize(data)
	if newString == "" {
		return "", errorHandler.ErrBadRequest
	}

	return newString, nil
}

type allowedModels interface {
	*models.Resume | *models.UserAccount | *models.Vacancy
}

func SanitizeObject[T allowedModels](obj T) (T, error) {
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		f := valueField.Interface()
		if valueStr, ok := f.(string); ok {
			sanitazeStr, err := SanitizeString(valueStr)
			if err != nil {
				return nil, err
			}
			valueField.SetString(sanitazeStr)
		}
	}
	return obj, nil
}
