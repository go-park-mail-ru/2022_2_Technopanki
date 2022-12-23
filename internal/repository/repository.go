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
	NotificationRepository    NotificationRepository
}

func NewPostgresRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:            impl.NewUserPostgres(db),
		ResumeRepository:          impl.NewResumePostgres(db),
		VacancyRepository:         impl.NewVacancyPostgres(db),
		VacancyActivityRepository: impl.NewVacancyActivityPostgres(db),
		NotificationRepository:    impl.NewNotificationPostgres(db),
	}
}

//go:generate mockgen -source repository.go -destination=mocks/mock.go

type UserRepository interface {
	CreateUser(user *models.UserAccount) error
	GetUserByEmail(email string) (*models.UserAccount, error)
	UpdateUser(newUser *models.UserAccount) error
	GetUser(id uint) (*models.UserAccount, error)
	GetAllUsers(conditions []string, filterValues []interface{}, flag string) ([]*models.UserAccount, error)
	GetUserSafety(id uint, safeFields []string) (*models.UserAccount, error)
	FindApplicantsToMailing() ([]string, error)
	FindNewVacancies() ([]*models.VacancyPreview, error)
	FindEmployersToMailing() ([]string, error)
	FindNewResumes() ([]*models.ResumePreview, error)
}

type VacancyRepository interface {
	GetAll(conditions []string, filterValues []interface{}) ([]*models.Vacancy, error)
	GetAllFilter(filter string) ([]*models.Vacancy, error)
	GetById(vacancyId uint) (*models.Vacancy, error)
	GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error)
	GetByUserId(userId uint) ([]*models.Vacancy, error)
	Create(vacancy *models.Vacancy) (uint, error)
	Update(userId uint, vacancyId uint, oldVacancy *models.Vacancy, updates *models.Vacancy) error
	Delete(userId, vacancyId uint) error
	AddVacancyToFavorites(user *models.UserAccount, vacancy *models.Vacancy) error
	GetUserFavoriteVacancies(user *models.UserAccount) ([]*models.Vacancy, error)
	DeleteVacancyFromFavorites(user *models.UserAccount, vacancy *models.Vacancy) error
	CheckFavoriteVacancy(userId uint, vacancyId uint) (bool, error)
}

type VacancyActivityRepository interface {
	ApplyForVacancy(*models.VacancyActivity) error
	GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivityPreview, error)
	GetAllUserApplies(userId uint) ([]*models.VacancyActivityPreview, error)
	DeleteUserApply(userId, applyId uint) error
}

type ResumeRepository interface {
	GetResume(id uint) (*models.Resume, error)
	GetAllResumes(conditions []string, filterValues []interface{}) ([]*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error)
	GetResumeInPDF(resumeId uint) (*models.ResumeInPDF, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
	GetEmployerIdByVacancyActivity(id uint) (uint, error)
}

type NotificationRepository interface {
	GetNotificationPreviewApply(id uint) (*models.NotificationPreview, error)
	GetNotificationPreviewDownloadPDF(id uint) (*models.NotificationPreview, error)
	GetApplyNotificationsByUser(id uint) ([]*models.NotificationPreview, error)
	GetDownloadPDFNotificationsByUser(id uint) ([]*models.NotificationPreview, error)
	CreateNotification(notification *models.Notification) error
	ReadNotification(id uint) error
	ReadAllNotifications(userId uint) error
	GetNotification(id uint) (*models.Notification, error)
	DeleteNotificationsFromUser(userId uint) error
}
