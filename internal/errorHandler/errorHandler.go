package errorHandler

import (
	"net/http"
)

var (
	ErrBadRequest         = newNonDescError("bad request")
	ErrUnauthorized       = newNonDescError("unauthorized")
	ErrServiceUnavailable = newNonDescError("service unavailable")
	ErrUserExists         = newSimpleDescError("Пользователь с таким email уже существует", "type", "email")
	ErrUserNotExists      = newNonDescError("user not found")
	ErrInvalidParam       = newNonDescError("invalid param")
	ErrSessionNotFound    = newNonDescError("session with this token not found")
	ErrVacancyNotFound    = newNonDescError("vacancy not found")
	ErrResumeNotFound     = newNonDescError("resume not found")

	ErrForbidden           = newNonDescError("forbidden")
	ErrWrongPassword       = newSimpleDescError("wrong password", "type", "password")
	ErrInvalidFileFormat   = newNonDescError("invalid file format")
	IncorrectNameLength    = newSimpleDescError("Длина имени должна быть между 3 и 20 символами", "type", "name")
	IncorrectSurnameLength = newSimpleDescError("Длина фамилии должна быть между 3 и 20 символами", "type", "surname")

	InvalidEmailFormat   = newSimpleDescError("email должен содержать @", "type", "email")
	IncorrectEmailLength = newSimpleDescError("Длина email должна быть между 8 and 30 символами", "type", "email")

	InvalidPasswordFormat   = newSimpleDescError("Пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)", "type", "password")
	IncorrectPasswordLength = newSimpleDescError("Длина пароля должна быть между 8 и 20 символами", "type", "password")

	InvalidUserType = newNonDescError("invalid input user type")
)

var errorToCode = map[error]int{
	ErrBadRequest:         http.StatusBadRequest,
	ErrUnauthorized:       http.StatusUnauthorized,
	ErrServiceUnavailable: http.StatusServiceUnavailable,
	ErrUserExists:         http.StatusBadRequest,
	ErrUserNotExists:      http.StatusUnauthorized,
	ErrInvalidParam:       http.StatusBadRequest,
	ErrSessionNotFound:    http.StatusUnauthorized,
	ErrVacancyNotFound:    http.StatusNotFound,
	ErrResumeNotFound:     http.StatusNotFound,

	ErrForbidden:           http.StatusForbidden,
	ErrWrongPassword:       http.StatusBadRequest,
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
