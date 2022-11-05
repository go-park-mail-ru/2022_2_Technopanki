package utils

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(user, input *models.UserAccount) error {
	if cryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); cryptErr != nil {
		return errorHandler.ErrWrongPassword
	}
	return nil
}

func GeneratePassword(password string, cfg *configs.CryptConfig) (string, error) {
	encryptedPassword, encryptErr := bcrypt.GenerateFromPassword([]byte(password), cfg.COST)
	if encryptErr != nil {
		return "", encryptErr
	}
	return string(encryptedPassword), nil
}
