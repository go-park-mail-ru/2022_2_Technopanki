package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func newUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func queryUserValidation(query *gorm.DB) error {
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errorHandler.ErrUserNotExists
	}
	return nil
}

func (up *UserPostgres) CreateUser(user *models.UserAccount) error {
	return up.db.Create(user).Error
}

func (up *UserPostgres) UpgradeUser(oldUser, newUser *models.UserAccount) error {
	return up.db.Model(oldUser).Updates(newUser).Error
}

func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Where("email = ?", email).Find(&result)
	return &result, queryUserValidation(query)
}

func (up *UserPostgres) IsUserExist(email string) (bool, error) {
	_, getErr := up.GetUserByEmail(email)
	if getErr == nil {
		return true, nil
	}
	if getErr == errorHandler.ErrUserNotExists {
		return false, nil
	}
	return false, getErr
}
