package repository

import (
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository    UserRepository
	VacancyRepository VacancyRepository
	ResumeRepository  ResumeRepository
}

func queryUserValidation(query *gorm.DB, object string) error {
	if query.Error != nil {
		return fmt.Errorf("postgre query error: %s", query.Error.Error())
	}
	if query.RowsAffected == 0 {
		switch object {
		case "user":
			return errorHandler.ErrUserNotExists
		case "vacancy":
			return errorHandler.ErrVacancyNotFound
		case "resume":
			return errorHandler.ErrResumeNotFound
		default:
			return fmt.Errorf("record not found: %s", object)
		}
	}
	return nil
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: newUserPostgres(db),
	}
}

type UserRepository interface {
	CreateUser(user *models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	IsUserExist(email string) (bool, error)
	UpdateUser(oldUser, newUser *models.UserAccount) error
	UpdateUserField(oldUser, newUser *models.UserAccount, field string) error
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
	Get()
	Create(entity.Resume)
	Update()
	Delete()
}
