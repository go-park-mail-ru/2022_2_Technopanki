package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"mime/multipart"
)

type UseCases struct {
	User    User
	Vacancy Vacancy
	Resume  Resume
}

func NewUseCases(repos *repository.Repository, session session.Repository, _cfg *configs.Config) *UseCases {
	return &UseCases{
		User:   newUserService(repos.UserRepository, session, _cfg),
		Resume: newResumeService(repos.ResumeRepository, _cfg),
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
	Get()
	Create(entity.Vacancy)
	Update()
	Delete()
}

type Resume interface {
	GetResume(id uint) (*models.Resume, error)
	GetResumeByApplicant(userId uint) ([]*models.Resume, error)
	CreateResume(resume *models.Resume, userId uint) error
	UpdateResume(id uint, resume *models.Resume) error
	DeleteResume(id uint) error
}
