package repository

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository    UserRepository
	VacancyRepository VacancyRepository
	ResumeRepository  ResumeRepository
	Cfg               *configs.Config
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
