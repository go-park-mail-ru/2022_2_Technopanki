package errorHandler

import "errors"

var errorMap = map[string]error{
	"the user is not found":                     errors.New("the user is not found"),
	"invalid signing method":                    errors.New("invalid signing method"),
	"token claims are not of type *tokenClaims": errors.New("token claims are not of type *tokenClaims"),
}
var unknown = errors.New("unknown error")

func ReturnError(errorName string) error {
	resultErr, ok := errorMap[errorName]
	if ok {
		return resultErr
	} else {
		return unknown
	}
}
func ReturnErrorCase[T any](errorName string) (T, error) {
	var emptyVal T
	return emptyVal, ReturnError(errorName)
}
