package escaping

import (
	"HeadHunter/internal/entity/models"
	"html"
	"reflect"
)

type allowedModels interface {
	*models.Resume | *models.UserAccount | *models.Vacancy
}

func EscapingObject[T allowedModels](obj T) T {
	val := reflect.ValueOf(obj).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		f := valueField.Interface()
		if valueStr, ok := f.(string); ok {
			escapeStr := html.EscapeString(valueStr)
			valueField.SetString(escapeStr)
		}
	}
	return obj
}
