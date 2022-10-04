package errorHandler

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrUserExists          = errors.New("пользователь с таким email уже существует")
	ErrUserNotExists       = errors.New("пользователя с таким email не существует")
	ErrInvalidQuery        = errors.New("invalid query")
	ErrSessionNotFound     = errors.New("session with this token not found")
	ErrCannotDeleteSession = errors.New("cannot delete session")

	IncorrectNameLength    = errors.New("длина имени должна быть между 3 и 20 символами")
	IncorrectSurnameLength = errors.New("длина фамилии должна быть между 3 и 20 символами")

	InvalidEmailFormat   = errors.New("email должен содержать @")
	IncorrectEmailLength = errors.New("длина email должна быть между 8 and 30 символами")

	InvalidPasswordFormat   = errors.New("пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)")
	IncorrectPasswordLength = errors.New("длина пароля должна быть между 8 и 20 символами")

	InvalidUserRole = errors.New("invalid input user role")
)

var errorToCode = map[error]int{
	ErrBadRequest:          http.StatusBadRequest,
	ErrUnauthorized:        http.StatusUnauthorized,
	ErrServiceUnavailable:  http.StatusServiceUnavailable,
	ErrUserExists:          http.StatusBadRequest,
	ErrUserNotExists:       http.StatusUnauthorized,
	ErrInvalidQuery:        http.StatusBadRequest,
	ErrSessionNotFound:     http.StatusUnauthorized,
	ErrCannotDeleteSession: http.StatusBadRequest,

	IncorrectNameLength:    http.StatusBadRequest,
	IncorrectSurnameLength: http.StatusBadRequest,
	InvalidUserRole:        http.StatusBadRequest,

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
