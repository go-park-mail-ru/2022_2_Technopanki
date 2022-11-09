package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository            UserRepository
	VacancyRepository         VacancyRepository
	VacancyActivityRepository VacancyActivityRepository
	ResumeRepository          ResumeRepository
}

func notFound(object string) error {
	switch object {
	case "user":
		return errorHandler.ErrUserNotExists
	case "vacancy":
		return errorHandler.ErrVacancyNotFound
	case "resume":
		return errorHandler.ErrResumeNotFound
	default:
		return fmt.Errorf("%s not found", object)
	}
}

func queryValidation(query *gorm.DB, object string) error {
	if query.Error != nil {
		if errors.Is(query.Error, fmt.Errorf("record not found")) {
			return notFound(object)
		}
		return fmt.Errorf("postgre query error: %s", query.Error.Error())
	}
	if query.RowsAffected == 0 {
		return notFound(object)
	}
	return nil
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:            newUserPostgres(db),
		ResumeRepository:          newResumePostgres(db),
		VacancyRepository:         newVacancyPostgres(db),
		VacancyActivityRepository: newVacancyActivityPostgres(db),
	}
}

//go:generate mockgen -source repository.go -destination=mocks/mock.go

type UserRepository interface {
	CreateUser(user *models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	IsUserExist(email string) (bool, error)
	UpdateUser(oldUser, newUser *models.UserAccount) error
	UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error
	GetUser(id uint) (*models.UserAccount, error)
	GetUserSafety(id uint, safeFields []string) (*models.UserAccount, error)
}

type VacancyRepository interface { //TODO Сделать репозиторий вакансий
	GetAll() ([]*models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]*models.Vacancy, error)
	Create(vacancy *models.Vacancy) (uint, error)
	Update(userId uint, vacancyId int, oldVacancy *models.Vacancy, updates *models.Vacancy) error
	Delete(uint, int) error
}

type VacancyActivityRepository interface {
	ApplyForVacancy(*models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]*models.VacancyActivity, error)
	GetAllUserApplies(int) ([]models.VacancyActivity, error)
}

type ResumeRepository interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.Resume, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
}
