package repository

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository    UserRepository
	VacancyRepository VacancyRepository
	ResumeRepository  ResumeRepository
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
		UserRepository:   newUserPostgres(db),
		ResumeRepository: newResumePostgres(db),
	}
}

type UserRepository interface {
	CreateUser(user *models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	IsUserExist(email string) (bool, error)
	UpdateUser(oldUser, newUser *models.UserAccount) error
	UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error
	GetUser(id uint) (*models.UserAccount, error)
	GetUserSafety(id uint, safeFields []string) (*models.UserAccount, error)
	UpdateUserImage()
}

type VacancyRepository interface { //TODO Сделать репозиторий вакансий
	Get()
	Create(entity.Vacancy)
	Update()
	Delete()
}

type ResumeRepository interface { //TODO Сделать репозиторий резюме
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
}
