package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository/impl"
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
		UserRepository:            impl.NewUserPostgres(db),
		ResumeRepository:          impl.NewResumePostgres(db),
		VacancyRepository:         impl.NewVacancyPostgres(db),
		VacancyActivityRepository: impl.NewVacancyActivityPostgres(db),
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

type VacancyRepository interface {
	GetAll() ([]*models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]*models.Vacancy, error)
	Create(vacancy *models.Vacancy) (uint, error)
	Update(userId uint, vacancyId int, oldVacancy *models.Vacancy, updates *models.Vacancy) error
	Delete(uint, int) error
	GetAuthor(string) (*models.UserAccount, error)
}

type VacancyActivityRepository interface {
	ApplyForVacancy(*models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]*models.VacancyActivity, error)
	GetAllUserApplies(int) ([]*models.VacancyActivity, error)
	GetAuthor(string) (*models.UserAccount, error)
	DeleteUserApply(uint, int) error
}

type ResumeRepository interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
	GetAuthor(email string) (*models.UserAccount, error)
}
