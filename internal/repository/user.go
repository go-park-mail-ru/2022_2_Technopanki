package repository

import (
	"HeadHunter/internal/entity/Models"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user Models.UserAccount) error {
	return nil
}

func (up *UserPostgres) GetUserByEmail(email string) (*Models.UserAccount, error) {
	return nil, nil
}
