package validation

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"regexp"
)

var passwordPattern = "^(?=.*[0-9])(?=.*[!#%^*$])(?=.*[a-zA-Z]).{8,20}"
var emailPattern = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"

func IsValidateAuthData(user entity.User) error {
	emailMatch, invalidEmail := regexp.Match(emailPattern, []byte(user.Email))
	passwordMatch, invalidPassword := regexp.Match(passwordPattern, []byte(user.Password))

	if invalidEmail != nil || invalidPassword != nil {
		return errorHandler.InvalidValidatePattern
	}

	if !emailMatch {
		return errorHandler.InvalidUserEmail
	}

	if !passwordMatch {
		return errorHandler.InvalidUserPassword
	}

	return nil
}
func IsValidate(user entity.User) error {
	if len([]rune(user.Name)) > 20 || len([]rune(user.Name)) < 3 {
		return errorHandler.InvalidUserName
	}

	if len([]rune(user.Surname)) > 20 || len([]rune(user.Surname)) < 3 {
		return errorHandler.InvalidUserSurname
	}

	if user.Role != "employer" && user.Role != "applicant" {
		return errorHandler.InvalidUserRole
	}

	return IsValidateAuthData(user)
}
