package repository

import (
	"HeadHunter/internal/entity/Models"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user Models.UserAccount) error {
	return up.db.Create(&user).Error
}

func (up *UserPostgres) GetUserByEmail(email string) (*Models.UserAccount, error) {
	var result Models.UserAccount
	query := up.db.Where("email = ?", email).Find(&result)
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			return nil, errorHandler.ErrUserNotExists
		}
		return nil, query.Error
	}
	return &result, nil
}
