package repository

import (
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByEmail(username string) (entity.User, error)
}

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
	}
}
