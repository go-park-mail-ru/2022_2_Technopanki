package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/escaping"
)

type VacancyService struct {
	vacancyRep repository.VacancyRepository
}

func NewVacancyService(vacancyRepos repository.VacancyRepository) *VacancyService {
	return &VacancyService{vacancyRep: vacancyRepos}
}

func (vs *VacancyService) GetAll() ([]*models.Vacancy, error) {
	return vs.vacancyRep.GetAll()
}

func (vs *VacancyService) Create(userId uint, input *models.Vacancy) (uint, error) {
	var sanitizeErr error
	input = escaping.EscapingObject[*models.Vacancy](input)
	if sanitizeErr != nil {
		return 0, sanitizeErr
	}

	input.PostedByUserId = userId
	return vs.vacancyRep.Create(input)
}

func (vs *VacancyService) GetById(vacancyID int) (*models.Vacancy, error) {
	return vs.vacancyRep.GetById(vacancyID)
}

func (vs *VacancyService) GetByUserId(userId int) ([]*models.Vacancy, error) {
	return vs.vacancyRep.GetByUserId(userId)
}

func (vs *VacancyService) Delete(userId uint, vacancyId int) error {
	return vs.vacancyRep.Delete(userId, vacancyId)
}

func (vs *VacancyService) Update(userId uint, vacancyId int, updates *models.Vacancy) error {

	var sanitizeErr error
	updates = escaping.EscapingObject[*models.Vacancy](updates)
	if sanitizeErr != nil {
		return sanitizeErr
	}
	oldVacancy, getErr := vs.vacancyRep.GetById(vacancyId)
	if getErr != nil {
		return getErr
	}
	return vs.vacancyRep.Update(userId, vacancyId, oldVacancy, updates)
}
