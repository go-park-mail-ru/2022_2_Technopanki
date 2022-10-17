package repository

import (
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository UserRepository
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
	}
}

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByEmail(username string) (entity.User, error)
}
