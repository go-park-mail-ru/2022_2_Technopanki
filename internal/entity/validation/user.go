package validation

import (
	"HeadHunter/internal/entity/models"
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
func IsAuthDataValid(user models.UserAccount) error {

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
func IsUserValid(user models.UserAccount) error {
	if user.UserType == "applicant" {
		if len([]rune(user.ApplicantName)) > 20 || len([]rune(user.ApplicantName)) < 3 {
			return errorHandler.IncorrectNameLength
		}

		if len([]rune(user.ApplicantSurname)) > 20 || len([]rune(user.ApplicantSurname)) < 3 {
			return errorHandler.IncorrectSurnameLength
		}
	} else if user.UserType == "employer" {
		if len([]rune(user.CompanyName)) > 20 || len([]rune(user.CompanyName)) < 2 {
			return errorHandler.IncorrectNameLength
		}
	} else {
		return errorHandler.InvalidUserType
	}
	return IsAuthDataValid(user)
}
