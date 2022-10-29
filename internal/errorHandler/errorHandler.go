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
	ErrCannotCreateUser   = errors.New("cannot create user")

	ErrCannotCreateSession = errors.New("cannot create session")
	ErrSessionNotFound     = errors.New("session with this token not found")
	ErrCannotDeleteSession = errors.New("cannot delete session")

	ErrVacancyNotFound = errors.New("vacancy not found")

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
	ErrCannotCreateUser:   http.StatusServiceUnavailable,

	ErrCannotCreateSession: http.StatusInternalServerError,
	ErrSessionNotFound:     http.StatusUnauthorized,
	ErrCannotDeleteSession: http.StatusInternalServerError,

	ErrVacancyNotFound: http.StatusNotFound,

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
