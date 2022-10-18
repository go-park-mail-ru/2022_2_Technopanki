package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/repository"
)

type UseCases struct {
	User    User
	Vacancy Vacancy
	Resume  Resume
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

type Vacancy interface { //TODO Сделать юзкейс вакансий
	Get()
	Create(entity.Vacancy)
	Update()
	Delete()
}

type Resume interface { //TODO Сделать юзкейс резюме
	Get()
	Create(entity.Resume)
	Update()
	Delete()
}
