package repository

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/Models"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository    UserRepository
	VacancyRepository VacancyRepository
	ResumeRepository  ResumeRepository
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserPostgres(db),
	}
}

type UserRepository interface {
	CreateUser(user Models.UserAccount) error
	GetUserByEmail(email string) (*Models.UserAccount, error)
}

type VacancyRepository interface { //TODO Сделать репозиторий вакансий
	Get()
	Create(entity.Vacancy)
	Update()
	Delete()
}

type ResumeRepository interface { //TODO Сделать репозиторий резюме
	Get()
	Create(entity.Resume)
	Update()
	Delete()
}
