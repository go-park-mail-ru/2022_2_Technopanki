package usecases

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
)

type VacancyActivityService struct {
	vacancyActivityRep repository.VacancyActivityRepository
}

func newVacancyActivityService(vacancyActivityRepos repository.VacancyActivityRepository) *VacancyActivityService {
	return &VacancyActivityService{vacancyActivityRep: vacancyActivityRepos}
}

func (vas *VacancyActivityService) GetAllVacancyApplies(vacancyId int) ([]*models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllVacancyApplies(vacancyId)
}

func (vas *VacancyActivityService) ApplyForVacancy(userId uint, input *models.VacancyActivity) error {
	input.UserAccountId = userId
	return vas.vacancyActivityRep.ApplyForVacancy(input)
}

func (vas *VacancyActivityService) GetAllUserApplies(userId int) ([]models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllUserApplies(userId)
}
