package repository

import (
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
