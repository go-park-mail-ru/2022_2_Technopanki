package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases/impl"
	"HeadHunter/internal/usecases/mail"
	"HeadHunter/internal/usecases/sender"
	"mime/multipart"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCases struct {
	User            User
	Vacancy         Vacancy
	VacancyActivity VacancyActivity
	Resume          Resume
}

func NewUseCases(repos *repository.Repository, session session.Repository,
	_sender sender.Sender, _cfg *configs.Config, _mail mail.Mail) *UseCases {
	return &UseCases{
		User:            impl.NewUserService(repos.UserRepository, session, _mail, _cfg),
		Resume:          impl.NewResumeService(repos.ResumeRepository, _cfg, repos.UserRepository),
		Vacancy:         impl.NewVacancyService(repos.VacancyRepository, repos.UserRepository),
		VacancyActivity: impl.NewVacancyActivityService(repos.VacancyActivityRepository, repos.UserRepository),
	}
}

type User interface {
	SignUp(input *models.UserAccount) (string, error)
	SignIn(input *models.UserAccount) (string, error)
	Logout(token string) error
	AuthCheck(email string) (*models.UserAccount, error)
	UpdateUser(input *models.UserAccount) error
	GetUser(id uint) (*models.UserAccount, error)
	GetUserId(email string) (uint, error)
	GetUserSafety(id uint) (*models.UserAccount, error)
	GetUserByEmail(email string) (*models.UserAccount, error)
	UploadUserImage(user *models.UserAccount, fileHeader *multipart.FileHeader) (string, error)
	DeleteUserImage(user *models.UserAccount) error
	ConfirmUser(code string, email string) (string, error)
}

type Vacancy interface {
	GetAll() ([]*models.Vacancy, error)
	GetById(vacancyId uint) (*models.Vacancy, error)
	GetByUserId(userId uint) ([]*models.Vacancy, error)
	Create(email string, input *models.Vacancy) (uint, error)
	Update(email string, vacancyId uint, updates *models.Vacancy) error
	Delete(email string, vacancyId uint) error
}

type VacancyActivity interface {
	ApplyForVacancy(email string, vacancyId uint, input *models.VacancyActivity) error
	GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error)
	GetAllUserApplies(userid uint) ([]*models.VacancyActivity, error)
	DeleteUserApply(email string, apply uint) error
}

type Resume interface {
	GetResume(id uint, email string) (*models.Resume, error)
	GetResumeByApplicant(userId uint, email string) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint, email string) ([]*models.ResumePreview, error)
	CreateResume(resume *models.Resume, email string) error
	UpdateResume(id uint, resume *models.Resume, email string) error
	DeleteResume(id uint, email string) error
}
