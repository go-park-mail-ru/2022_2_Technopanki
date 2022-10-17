package usecases

import (
	"HeadHunter/internal/entity"
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

type User interface {
	SignUp(input entity.User) (string, error)
	SignIn(input *entity.User) (string, error)
	Logout(token string) error
	AuthCheck(email string) (entity.User, error)
}
