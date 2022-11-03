package errorHandler

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest         = errors.New("bad request")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrUserExists         = errors.New("Пользователь с таким email уже существует")
	ErrUserNotExists      = errors.New("Пользователя с таким email не существует")
	ErrInvalidQuery       = errors.New("invalid query")
	ErrInvalidParam       = errors.New("invalid param")
	ErrSessionNotFound    = errors.New("session with this token not found")
	ErrVacancyNotFound    = errors.New("vacancy not found")
	ErrResumeNotFound     = errors.New("resume not found")

	ErrForbidden           = errors.New("forbidden")
	ErrWrongAnswer         = errors.New("wrong answer")
	ErrInvalidFileFormat   = errors.New("invalid file format")
	IncorrectNameLength    = errors.New("Длина имени должна быть между 3 и 20 символами")
	IncorrectSurnameLength = errors.New("Длина фамилии должна быть между 3 и 20 символами")

	InvalidEmailFormat   = errors.New("email должен содержать @")
	IncorrectEmailLength = errors.New("Длина email должна быть между 8 and 30 символами")

	InvalidPasswordFormat   = errors.New("Пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)")
	IncorrectPasswordLength = errors.New("Длина пароля должна быть между 8 и 20 символами")

	InvalidUserType = errors.New("invalid input user type")
)

var errorToCode = map[error]int{
	ErrBadRequest:         http.StatusBadRequest,
	ErrUnauthorized:       http.StatusUnauthorized,
	ErrServiceUnavailable: http.StatusServiceUnavailable,
	ErrUserExists:         http.StatusBadRequest,
	ErrUserNotExists:      http.StatusUnauthorized,
	ErrInvalidQuery:       http.StatusBadRequest,
	ErrInvalidParam:       http.StatusBadRequest,
	ErrSessionNotFound:    http.StatusUnauthorized,
	ErrVacancyNotFound:    http.StatusNotFound,
	ErrResumeNotFound:     http.StatusNotFound,

	ErrForbidden:           http.StatusForbidden,
	ErrWrongAnswer:         http.StatusBadRequest,
	ErrInvalidFileFormat:   http.StatusBadRequest,
	IncorrectNameLength:    http.StatusBadRequest,
	IncorrectSurnameLength: http.StatusBadRequest,
	InvalidUserType:        http.StatusBadRequest,

	InvalidEmailFormat:   http.StatusBadRequest,
	IncorrectEmailLength: http.StatusBadRequest,

	InvalidPasswordFormat:   http.StatusBadRequest,
	IncorrectPasswordLength: http.StatusBadRequest,
}

func ConvertError(err error) int {
	result, ok := errorToCode[err]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}
