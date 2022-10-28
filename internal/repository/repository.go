package repository

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
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
	CreateUser(user models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	IsUserExist(email string) (bool, error)
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
