package errorHandler

import (
	"net/http"
)

var (
	ErrBadRequest         = newComplexError("bad request")
	ErrUnauthorized       = newComplexError("unauthorized")
	ErrServiceUnavailable = newComplexError("service unavailable")
	ErrUserExists         = newComplexError("Пользователь с таким email уже существует", "email")
	ErrUserNotExists      = newComplexError("user not found")
	ErrInvalidParam       = newComplexError("invalid param")
	ErrSessionNotFound    = newComplexError("session with this token not found")
	ErrVacancyNotFound    = newComplexError("vacancy not found")
	ErrResumeNotFound     = newComplexError("resume not found")

	ErrForbidden           = newComplexError("forbidden")
	ErrWrongPassword       = newComplexError("wrong password", "password")
	ErrInvalidFileFormat   = newComplexError("invalid file format")
	IncorrectNameLength    = newComplexError("Длина имени должна быть между 2 и 30 символами", "name")
	IncorrectSurnameLength = newComplexError("Длина фамилии должна быть между 2 и 30 символами", "surname")

	InvalidEmailFormat   = newComplexError("email должен содержать @", "email")
	IncorrectEmailLength = newComplexError("Длина email должна быть между 8 and 30 символами", "email")

	InvalidPasswordFormat   = newComplexError("Пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)", "password")
	IncorrectPasswordLength = newComplexError("Длина пароля должна быть между 8 и 20 символами", "password")

	InvalidUserType = newComplexError("invalid input user type")
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
