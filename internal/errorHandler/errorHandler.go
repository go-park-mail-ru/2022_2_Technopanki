package errorHandler

import (
	"errors"
	"net/http"
)

var (
	ErrResumeNotFound          = errors.New("resume not found")
	ErrBadRequest              = errors.New("bad request")
	ErrUnauthorized            = errors.New("unauthorized")
	ErrServiceUnavailable      = errors.New("service unavailable")
	ErrUserExists              = errors.New("Пользователь с таким email уже существует")
	ErrUserNotExists           = errors.New("Пользователя с таким email не существует")
	ErrInvalidParam            = errors.New("invalid parameter")
	ErrCannotCreateUser        = errors.New("cannot create user")
	ErrCannotDeleteVacancy     = errors.New("cannot delete vacancy")
	ErrCannotUpdateVacancy     = errors.New("cannot update vacancy")
	ErrCannotApplyForVacancy   = errors.New("cannot apply fo vacancy")
	ErrUpdateStructHasNoValues = errors.New("update structure has no values")
	ErrCSRFTokenMismatched     = errors.New("csrf token mismatched")

	ErrCannotCreateSession = errors.New("cannot create session")
	ErrSessionNotFound     = errors.New("session with this token not found")
	ErrCannotDeleteSession = errors.New("cannot delete session")

	ErrVacancyNotFound     = errors.New("vacancy not found")
	ErrCannotDeleteAvatar  = errors.New("Невозможно удалить аватар")
	ErrForbidden           = errors.New("forbidden")
	ErrWrongPassword       = errors.New("wrong password")
	ErrInvalidFileFormat   = errors.New("invalid file format")
	IncorrectNameLength    = errors.New("Длина имени должна быть между 2 и 30 символами")
	IncorrectSurnameLength = errors.New("Длина фамилии должна быть между 2 и 30 символами")

	InvalidEmailFormat   = errors.New("email должен содержать @")
	IncorrectEmailLength = errors.New("Длина email должна быть между 8 and 30 символами")

	InvalidPasswordFormat   = errors.New("Пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)")
	IncorrectPasswordLength = errors.New("Длина пароля должна быть между 8 и 20 символами")

	CSRFTokenMismatch = errors.New("CSRF token mismatch")
	InvalidUserType   = errors.New("invalid input user type")
)

var errorToCode = map[error]int{
	ErrResumeNotFound:        http.StatusNotFound,
	ErrBadRequest:            http.StatusBadRequest,
	ErrUnauthorized:          http.StatusUnauthorized,
	ErrServiceUnavailable:    http.StatusServiceUnavailable,
	ErrUserExists:            http.StatusBadRequest,
	ErrUserNotExists:         http.StatusUnauthorized,
	ErrInvalidParam:          http.StatusBadRequest,
	ErrCannotCreateUser:      http.StatusServiceUnavailable,
	ErrCannotDeleteVacancy:   http.StatusServiceUnavailable,
	ErrCannotUpdateVacancy:   http.StatusServiceUnavailable,
	ErrCannotCreateSession:   http.StatusInternalServerError,
	ErrSessionNotFound:       http.StatusUnauthorized,
	ErrCannotDeleteSession:   http.StatusInternalServerError,
	ErrCannotDeleteAvatar:    http.StatusBadRequest,
	ErrCannotApplyForVacancy: http.StatusBadRequest,

	ErrCSRFTokenMismatched: http.StatusForbidden,

	ErrVacancyNotFound:         http.StatusNotFound,
	ErrUpdateStructHasNoValues: http.StatusInternalServerError,

	ErrForbidden:           http.StatusForbidden,
	ErrWrongPassword:       http.StatusBadRequest,
	ErrInvalidFileFormat:   http.StatusBadRequest,
	IncorrectNameLength:    http.StatusBadRequest,
	IncorrectSurnameLength: http.StatusBadRequest,
	CSRFTokenMismatch:      http.StatusBadRequest,
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

var errorDescriptors = map[error][]string{
	ErrUserExists:           {"email"},
	ErrUserNotExists:        {"email"},
	ErrWrongPassword:        {"password"},
	IncorrectNameLength:     {"name"},
	IncorrectSurnameLength:  {"surname"},
	InvalidEmailFormat:      {"email"},
	IncorrectEmailLength:    {"email"},
	InvalidPasswordFormat:   {"password"},
	IncorrectPasswordLength: {"password"},
}

func GetErrorDescriptors(err error) []string {
	result, ok := errorDescriptors[err]
	if ok {
		return result
	}
	return []string{}
}
