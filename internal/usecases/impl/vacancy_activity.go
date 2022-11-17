package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
)

type VacancyActivityService struct {
	vacancyActivityRep repository.VacancyActivityRepository
}

func NewVacancyActivityService(vacancyActivityRepos repository.VacancyActivityRepository) *VacancyActivityService {
	return &VacancyActivityService{vacancyActivityRep: vacancyActivityRepos}
}

func (vas *VacancyActivityService) ApplyForVacancy(userId uint, vacancyId uint, input *models.VacancyActivity) error {
	input.UserAccountId = userId
	input.VacancyId = vacancyId
	return vas.vacancyActivityRep.ApplyForVacancy(input)
}

func (vas *VacancyActivityService) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllVacancyApplies(vacancyId)
}

func (vas *VacancyActivityService) GetAllUserApplies(userId uint) ([]*models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllUserApplies(userId)
}
