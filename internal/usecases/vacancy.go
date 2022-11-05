package usecases

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/repository"
	"strconv"
)

type VacancyService struct {
	vacancyRep repository.VacancyRepository
}

func newVacancyService(vacancyRepos repository.VacancyRepository) *VacancyService {
	return &VacancyService{vacancyRep: vacancyRepos}
}

func (vs *VacancyService) GetAll() ([]models.Vacancy, error) {
	return vs.vacancyRep.GetAll()
}

func (vs *VacancyService) Create(userId uint, input *models.Vacancy) (uint, error) {
	input.PostedByUserId = userId
	return vs.vacancyRep.Create(input)
}

func (vs *VacancyService) GetById(vacancyID int) (*models.Vacancy, error) {
	return vs.vacancyRep.GetById(vacancyID)
}

func (vs *VacancyService) GetByUserId(userId int) ([]models.Vacancy, error) {
	return vs.vacancyRep.GetByUserId(userId)
}

func (vs *VacancyService) Delete(userId uint, vacancyId int) error {
	return vs.vacancyRep.Delete(userId, vacancyId)
}

func (vs *VacancyService) Update(userId uint, vacancyId int, updates *models.UpdateVacancy) error {
	userIdString := strconv.FormatUint(uint64(userId), 10)
	vacancyIdString := strconv.Itoa(vacancyId)
	if err := validation.UpdateVacancyValidate(*updates); err != nil {
		return err
	}
	return vs.vacancyRep.Update(userIdString, vacancyIdString, updates)
}
