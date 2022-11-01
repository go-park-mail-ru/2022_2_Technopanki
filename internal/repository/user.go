package repository

import (
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user models.UserAccount) error {
	return nil
}

func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	return nil, nil
}
