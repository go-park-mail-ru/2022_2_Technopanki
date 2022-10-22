package usecases

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/Models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
)

type UseCases struct {
	User    User
	Vacancy Vacancy
	Resume  Resume
}

func NewUseCases(repos *repository.Repository, session session.Repository) *UseCases {
	return &UseCases{
		User: newUserService(repos.UserRepository, session),
	}
}

type User interface {
	SignUp(input Models.UserAccount) (string, error)
	SignIn(input *Models.UserAccount) (string, error)
	Logout(token string) error
	AuthCheck(email string) (Models.UserAccount, error)
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
