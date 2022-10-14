package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/repository"
)

type User interface {
	CreateUser(user entity.User) error
	GetUserByEmail(username string) (entity.User, error)
}

type UseCases struct {
	User User
}

func newUseCases(repos *repository.Repository) *UseCases {
	return &UseCases{
		User: newUserService(repos.UserRepository),
	}
}
