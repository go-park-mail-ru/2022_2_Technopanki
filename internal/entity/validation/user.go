package validation

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/errorHandler"
	"strings"
)

func verifyPassword(password string) bool {
	var special, number, symbol bool
	for _, c := range password {
		if c >= '0' && c <= '9' {
			number = true
		} else if strings.Contains("!#%^$", string(c)) {
			special = true
		} else if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			symbol = true
		} else {
			return false
		}
	}
	return number && special && symbol
}
func IsAuthDataValid(user entity.User) error {

	if strings.Count(user.Email, "@") != 1 {
		return errorHandler.InvalidEmailFormat
	}

	if len(user.Email) < 8 || len(user.Email) > 30 {
		return errorHandler.IncorrectEmailLength
	}

	if !verifyPassword(user.Password) {
		return errorHandler.InvalidPasswordFormat
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return errorHandler.IncorrectPasswordLength
	}

	return nil
}
func IsUserValid(user entity.User) error {
	if len([]rune(user.Name)) > 20 || len([]rune(user.Name)) < 3 {
		return errorHandler.IncorrectNameLength
	}

	if len([]rune(user.Surname)) > 20 || len([]rune(user.Surname)) < 3 {
		return errorHandler.IncorrectSurnameLength
	}

	if user.Role != "employer" && user.Role != "applicant" {
		return errorHandler.InvalidUserRole
	}

	return IsAuthDataValid(user)
}
