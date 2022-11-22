package validation

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/utils"
	"HeadHunter/pkg/errorHandler"
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

func IsAuthDataValid(user *models.UserAccount, cfg configs.ValidationConfig) error {

	if strings.Count(user.Email, "@") != 1 {
		return errorHandler.InvalidEmailFormat
	}

	if len(user.Email) < cfg.MinEmailLength || len(user.Email) > cfg.MaxEmailLength {
		return errorHandler.IncorrectEmailLength
	}

	if !verifyPassword(user.Password) {
		return errorHandler.InvalidPasswordFormat
	}

	if len(user.Password) < cfg.MinPasswordLength || len(user.Password) > cfg.MaxPasswordLength {
		return errorHandler.IncorrectPasswordLength
	}

	return nil
}

func IsMainDataValid(user *models.UserAccount, cfg configs.ValidationConfig) error {
	if err := AllowedFieldsValidation(user); err != nil {
		return err
	}

	if user.UserType == "applicant" {
		if len([]rune(user.ApplicantName)) > cfg.MaxNameLength || len([]rune(user.ApplicantName)) < cfg.MinNameLength {
			return errorHandler.IncorrectNameLength
		}

		if len([]rune(user.ApplicantSurname)) > cfg.MaxSurnameLength || len([]rune(user.ApplicantSurname)) < cfg.MinSurnameLength {
			return errorHandler.IncorrectSurnameLength
		}
	} else if user.UserType == "employer" {
		if len([]rune(user.CompanyName)) > cfg.MaxNameLength || len([]rune(user.CompanyName)) < cfg.MinNameLength {
			return errorHandler.IncorrectNameLength
		}
	} else {
		return errorHandler.InvalidUserType
	}
	return nil
}

func IsUserValid(user *models.UserAccount, cfg configs.ValidationConfig) error {
	if mainDataErr := IsMainDataValid(user, cfg); mainDataErr != nil {
		return mainDataErr
	}
	return IsAuthDataValid(user, cfg)
}

func AllowedFieldsValidation(user *models.UserAccount) error {
	fields := strings.Split(user.PublicFields, " ")
	if len(fields) > len(models.PrivateUserFields) {
		return errorHandler.ErrBadRequest
	}

	if len(fields) == 1 && (fields[0] == "" || fields[0] == models.NoPublicFields) {
		return nil
	}

	for _, elem := range fields {
		if !utils.HasStringArrayElement(elem, models.PrivateUserFields) {
			return errorHandler.ErrBadRequest
		}
	}
	return nil
}
