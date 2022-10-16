package usecases

import (
	"HeadHunter/internal/repository"
)

type UseCases struct {
	User User
}

func NewUseCases(repos *repository.Repository) *UseCases {
	return &UseCases{
		User: newUserService(repos.UserRepository),
	}
}
