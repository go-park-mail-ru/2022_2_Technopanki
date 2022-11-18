package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases/impl"
	"mime/multipart"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCases struct {
	User            User
	Vacancy         Vacancy
	VacancyActivity VacancyActivity
	Resume          Resume
}

func NewUseCases(repos *repository.Repository, session session.Repository, _cfg *configs.Config) *UseCases {
	return &UseCases{
		User:            impl.NewUserService(repos.UserRepository, session, _cfg),
		Resume:          impl.NewResumeService(repos.ResumeRepository, _cfg),
		Vacancy:         impl.NewVacancyService(repos.VacancyRepository),
		VacancyActivity: impl.NewVacancyActivityService(repos.VacancyActivityRepository),
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
}

type Vacancy interface {
	GetAll() ([]*models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]*models.Vacancy, error)
	Create(string, *models.Vacancy) (uint, error)
	Update(string, int, *models.Vacancy) error
	Delete(string, int) error
}

type VacancyActivity interface {
	ApplyForVacancy(string, int, *models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]*models.VacancyActivity, error)
	GetAllUserApplies(int) ([]*models.VacancyActivity, error)
	DeleteUserApply(string, int) error
}

type Resume interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint, email string) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint, email string) ([]*models.ResumePreview, error)
	CreateResume(resume *models.Resume, email string) error
	UpdateResume(id uint, resume *models.Resume, email string) error
	DeleteResume(id uint, email string) error
}
