package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"fmt"
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
		if query.Error.Error() == "record not found" {
			return nil, errorHandler.ErrUserNotExists
		}
		return nil, query.Error
	}
	fmt.Println(query.RowsAffected)
	if query.RowsAffected == 0 {
		return nil, errorHandler.ErrUserNotExists
	}
	fmt.Println(result)
	return &result, nil
}

func (up *UserPostgres) IsUserExist(email string) (bool, error) {
	return true, nil
}
