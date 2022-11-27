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
	UpdateUser(oldUser, newUser *models.UserAccount) error
	UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error
	GetUser(id uint) (*models.UserAccount, error)
	GetUserSafety(id uint, safeFields []string) (*models.UserAccount, error)
	UpdatePassword(user *models.UserAccount) error
	GetBestVacanciesForApplicant(user *models.UserAccount) ([]*models.Vacancy, error)
	GetBestApplicantForEmployer(user *models.UserAccount) ([]*models.UserAccount, error)
}

type VacancyRepository interface {
	GetAll() ([]*models.Vacancy, error)
	GetById(vacancyId uint) (*models.Vacancy, error)
	GetByUserId(userId uint) ([]*models.Vacancy, error)
	Create(vacancy *models.Vacancy) (uint, error)
	Update(userId uint, vacancyId uint, oldVacancy *models.Vacancy, updates *models.Vacancy) error
	Delete(userId, vacancyId uint) error
}

type VacancyActivityRepository interface {
	ApplyForVacancy(*models.VacancyActivity) error
	GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error)
	GetAllUserApplies(userId uint) ([]*models.VacancyActivity, error)
	DeleteUserApply(userId, applyId uint) error
}

type ResumeRepository interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
	GetEmployerIdByVacancyActivity(id uint) (uint, error)
}
