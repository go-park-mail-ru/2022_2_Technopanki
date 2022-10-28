package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user models.UserAccount) error {
	return up.db.Create(&user).Error
}

func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	var count int64
	query := up.db.Where("email = ?", email).Find(&result).Count(&count)
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			return nil, errorHandler.ErrUserNotExists
		}
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errorHandler.ErrUserNotExists
	}
	return &result, nil
}
