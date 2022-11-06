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
	ErrBadRequest.err:         http.StatusBadRequest,
	ErrUnauthorized.err:       http.StatusUnauthorized,
	ErrServiceUnavailable.err: http.StatusServiceUnavailable,
	ErrUserExists.err:         http.StatusBadRequest,
	ErrUserNotExists.err:      http.StatusUnauthorized,
	ErrInvalidParam.err:       http.StatusBadRequest,
	ErrSessionNotFound.err:    http.StatusUnauthorized,
	ErrVacancyNotFound.err:    http.StatusNotFound,
	ErrResumeNotFound.err:     http.StatusNotFound,

	ErrForbidden.err:           http.StatusForbidden,
	ErrWrongPassword.err:       http.StatusBadRequest,
	ErrInvalidFileFormat.err:   http.StatusBadRequest,
	IncorrectNameLength.err:    http.StatusBadRequest,
	IncorrectSurnameLength.err: http.StatusBadRequest,
	InvalidUserType.err:        http.StatusBadRequest,

	InvalidEmailFormat.err:   http.StatusBadRequest,
	IncorrectEmailLength.err: http.StatusBadRequest,

	InvalidPasswordFormat.err:   http.StatusBadRequest,
	IncorrectPasswordLength.err: http.StatusBadRequest,
}

func ConvertError(err error) int {
	complexErr, ok := err.(*ComplexError)
	if ok {
		return ConvertError(complexErr.err)
	}
	result, ok := errorToCode[err]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}
