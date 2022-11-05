package repository

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository            UserRepository
	VacancyRepository         VacancyRepository
	VacancyActivityRepository VacancyActivityRepository
	ResumeRepository          ResumeRepository
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:            newUserPostgres(db),
		VacancyRepository:         newVacancyPostgres(db),
		VacancyActivityRepository: newVacancyActivityPostgres(db),
	}
}

type UserRepository interface {
	CreateUser(user *models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	IsUserExist(email string) (bool, error)
	UpgradeUser(oldUser, newUser *models.UserAccount) error
}

type VacancyRepository interface { //TODO Сделать репозиторий вакансий
	GetAll() ([]models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]models.Vacancy, error)
	Create(vacancy *models.Vacancy) (uint, error)
	Update(string, string, *models.UpdateVacancy) error
	Delete(uint, int) error
}

type VacancyActivityRepository interface {
	ApplyForVacancy(*models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]models.VacancyActivity, error)
	GetAllUserApplies(int) ([]models.VacancyActivity, error)
}

type ResumeRepository interface { //TODO Сделать репозиторий резюме
	Get()
	Create(entity.Resume)
	Update()
	Delete()
}
