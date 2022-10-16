package repository

import (
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByEmail(username string) (entity.User, error)
}

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user entity.User) error {
	return nil
}

func (up *UserPostgres) GetUserByEmail(username string) (entity.User, error) {
	return entity.User{}, nil
}
