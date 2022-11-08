package usecases

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
)

type UseCases struct {
	User            User
	Vacancy         Vacancy
	VacancyActivity VacancyActivity
	Resume          Resume
}

func NewUseCases(repos *repository.Repository, session session.Repository, _cfg *configs.Config) *UseCases {
	return &UseCases{
		User:            newUserService(repos.UserRepository, session, _cfg),
		Vacancy:         newVacancyService(repos.VacancyRepository),
		VacancyActivity: newVacancyActivityService(repos.VacancyActivityRepository),
	}
}

type User interface {
	SignUp(input *models.UserAccount) (string, error)
	SignIn(input *models.UserAccount) (string, error)
	Logout(token string) error
	AuthCheck(email string) (*models.UserAccount, error)
	UpgradeUser(input *models.UserAccount) error
	GetUserId(email string) (uint, error)
}

type Vacancy interface { //TODO Сделать юзкейс вакансий
	GetAll() ([]*models.Vacancy, error)
	GetById(int) (*models.Vacancy, error)
	GetByUserId(int) ([]models.Vacancy, error)
	Create(uint, *models.Vacancy) (uint, error)
	Update(uint, int, *models.Vacancy) error
	Delete(uint, int) error
}

type VacancyActivity interface {
	ApplyForVacancy(uint, *models.VacancyActivity) error
	GetAllVacancyApplies(int) ([]*models.VacancyActivity, error)
	GetAllUserApplies(int) ([]models.VacancyActivity, error)
}

type Resume interface { //TODO Сделать юзкейс резюме
	Get()
	Create(entity.Resume)
	Update()
	Delete()
}
