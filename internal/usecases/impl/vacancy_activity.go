package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"HeadHunter/internal/repository"
)

type VacancyActivityService struct {
	vacancyActivityRep repository.VacancyActivityRepository
}

func NewVacancyActivityService(vacancyActivityRepos repository.VacancyActivityRepository) *VacancyActivityService {
	return &VacancyActivityService{vacancyActivityRep: vacancyActivityRepos}
}

func (vas *VacancyActivityService) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllVacancyApplies(vacancyId)
}

func (vas *VacancyActivityService) ApplyForVacancy(email string, vacancyId uint, input *models.VacancyActivity) error {
	user, getErr := vas.vacancyActivityRep.GetAuthor(email)
	if getErr != nil {
		return getErr
	}
	if user.UserType != "applicant" {
		return errorHandler.InvalidUserType
	}

	input.UserAccountId = user.ID
	input.VacancyId = uint(vacancyId)
	return vas.vacancyActivityRep.ApplyForVacancy(input)
}

func (vas *VacancyActivityService) GetAllUserApplies(userId uint) ([]*models.VacancyActivity, error) {
	return vas.vacancyActivityRep.GetAllUserApplies(userId)
}

func (vas *VacancyActivityService) DeleteUserApply(email string, id uint) error {
	user, getErr := vas.vacancyActivityRep.GetAuthor(email)
	if getErr != nil {
		return getErr
	}
	userId := user.ID
	return vas.vacancyActivityRep.DeleteUserApply(userId, id)
}
