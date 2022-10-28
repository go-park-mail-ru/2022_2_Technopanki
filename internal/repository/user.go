package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//COST The cost of the password encryption algorithm
var COST = 10

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user models.UserAccount) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), COST)
	if err != nil {
		return err
	}
	user.Password = string(encryptedPassword)
	return up.db.Create(&user).Error
}

func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Where("email = ?", email).Find(&result)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errorHandler.ErrUserNotExists
	}
	return &result, nil
}
