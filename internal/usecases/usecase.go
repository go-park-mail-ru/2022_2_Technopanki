package usecases

import (
	"HeadHunter/common/session"
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/impl"
	"HeadHunter/internal/usecases/mail"
	"mime/multipart"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCases struct {
	User            User
	Vacancy         Vacancy
	VacancyActivity VacancyActivity
	Resume          Resume
	Notification    Notification
	Mail            mail.Mail
}

func NewUseCases(repos *repository.Repository, session session.Repository, _cfg *configs.Config, _mail mail.Mail) *UseCases {
	return &UseCases{
		User:            impl.NewUserService(repos.UserRepository, session, _mail, _cfg),
		Resume:          impl.NewResumeService(repos.ResumeRepository, _cfg, repos.UserRepository),
		Vacancy:         impl.NewVacancyService(repos.VacancyRepository, repos.UserRepository),
		VacancyActivity: impl.NewVacancyActivityService(repos.VacancyActivityRepository, repos.UserRepository, repos.VacancyRepository, repos.NotificationRepository),
		Notification:    impl.NewNotificationService(repos.NotificationRepository, repos.UserRepository),
		Mail:            _mail,
	}
}

type User interface {
	SignUp(input *models.UserAccount) (string, error)
	SignIn(input *models.UserAccount) (string, error)
	Logout(token string) error
	AuthCheck(email string) (*models.UserAccount, error)
	UpdateUser(input *models.UserAccount) error
	GetUser(id uint) (*models.UserAccount, error)
	GetAllEmployers(filters models.UserFilter) ([]*models.UserAccount, error)
	GetAllApplicants(filters models.UserFilter) ([]*models.UserAccount, error)
	GetUserId(email string) (uint, error)
	GetUserSafety(id uint) (*models.UserAccount, error)
	GetUserByEmail(email string) (*models.UserAccount, error)
	UploadUserImage(user *models.UserAccount, fileHeader *multipart.FileHeader) (string, error)
	DeleteUserImage(user *models.UserAccount) error
	ConfirmUser(code, email string) (*models.UserAccount, string, error)
	UpdatePassword(code, email, password string) error
	GetMailing(email string) error
}

type Vacancy interface {
	GetAll(filters models.VacancyFilter) ([]*models.Vacancy, error)
	GetById(vacancyId uint) (*models.Vacancy, error)
	GetByUserId(userId uint) ([]*models.Vacancy, error)
	GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error)
	Create(email string, input *models.Vacancy) (uint, error)
	Update(email string, vacancyId uint, updates *models.Vacancy) error
	Delete(email string, vacancyId uint) error
	AddVacancyToFavorites(email string, vacancyId uint) error
	GetUserFavoriteVacancies(email string) ([]*models.Vacancy, error)
	DeleteVacancyFromFavorites(email string, vacancyId uint) error
	CheckFavoriteVacancy(email string, vacancyId uint) (bool, error)
}

type VacancyActivity interface {
	ApplyForVacancy(email string, vacancyId uint, input *models.VacancyActivity) (*models.Notification, error)
	GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivityPreview, error)
	GetAllUserApplies(userid uint) ([]*models.VacancyActivityPreview, error)
	DeleteUserApply(email string, apply uint) error
}

type Resume interface {
	GetResume(id uint) (*models.Resume, error)
	GetAllResumes(filters models.ResumeFilter) ([]*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error)
	GetResumeInPDF(email string, resumeId uint, style string) (*models.Notification, []byte, error)
	CreateResume(resume *models.Resume, email string) error
	UpdateResume(id uint, resume *models.Resume, email string) error
	DeleteResume(id uint, email string) error
}

type Notification interface {
	GetNotificationsByEmail(email string) ([]*models.NotificationPreview, error)
	CreateNotification(notification *models.Notification) (*models.NotificationPreview, error)
	ReadNotification(email string, id uint) error
	ReadAllNotifications(email string) error
	ClearNotifications(email string) error
}
