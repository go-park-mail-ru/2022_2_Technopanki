package usecases

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository/impl"
	"gorm.io/gorm"
)

func GetUser(email string, db *gorm.DB) (*models.UserAccount, error) {
	var result models.UserAccount
	query := db.Where("email = ?", email).Find(&result)
	return &result, impl.QueryValidation(query, "user")
}
