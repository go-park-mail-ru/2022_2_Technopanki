package errorHandler

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrUserExists          = errors.New("user already exists")
	ErrUserNotExists       = errors.New("user not exists")
	ErrInvalidQuery        = errors.New("invalid query")
	ErrNoSuitableSession   = errors.New("no session with this token")
	ErrCannotDeleteSession = errors.New("cannot delete session")
)

var errorToCode = map[error]int{
	ErrBadRequest:          http.StatusBadRequest,
	ErrUnauthorized:        http.StatusUnauthorized,
	ErrServiceUnavailable:  http.StatusServiceUnavailable,
	ErrUserExists:          http.StatusBadRequest,
	ErrUserNotExists:       http.StatusUnauthorized,
	ErrInvalidQuery:        http.StatusBadRequest,
	ErrNoSuitableSession:   http.StatusUnauthorized,
	ErrCannotDeleteSession: http.StatusBadRequest,
}

func ConvertError(err error) int {
	result, ok := errorToCode[err]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}
