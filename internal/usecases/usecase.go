package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
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
		User:            newUserService(repos.UserRepository, session, _cfg),
		Resume:          newResumeService(repos.ResumeRepository, _cfg),
		Vacancy:         newVacancyService(repos.VacancyRepository),
		VacancyActivity: newVacancyActivityService(repos.VacancyActivityRepository),
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

type Vacancy interface { //TODO Сделать юзкейс вакансий
	GetAll() ([]*models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]*models.Vacancy, error)
	Create(uint, *models.Vacancy) (uint, error)
	Update(uint, int, *models.Vacancy) error
	Delete(uint, int) error
}

type VacancyActivity interface {
	ApplyForVacancy(uint, int, *models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]*models.VacancyActivity, error)
	GetAllUserApplies(int) ([]*models.VacancyActivity, error)
}

type Resume interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	GetPreviewResumeByApplicant(userId uint) ([]*models.Resume, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
}
