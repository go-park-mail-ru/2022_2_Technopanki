package repository

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository UserRepository
	Cfg            *configs.Config
}

func NewPostgresRepository(db *gorm.DB, _cfg *configs.Config) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
		Cfg:            _cfg,
	}
}

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByEmail(username string) (entity.User, error)
}
