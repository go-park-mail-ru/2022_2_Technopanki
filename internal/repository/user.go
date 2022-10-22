package repository

import (
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user entity.User) error {
	return nil
}

func (up *UserPostgres) GetUserByEmail(email string) (entity.User, error) {
	//var user = entity.User{Email: email}
	//up.db.First(&user)
	return entity.User{}, nil
}
