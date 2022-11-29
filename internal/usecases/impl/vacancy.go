package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/pkg/errorHandler"
	"reflect"
	"errors"
)

type VacancyService struct {
	vacancyRep repository.VacancyRepository
	userRep    repository.UserRepository
}

func NewVacancyService(vacancyRepos repository.VacancyRepository, _userRep repository.UserRepository) *VacancyService {
	return &VacancyService{vacancyRep: vacancyRepos, userRep: _userRep}
}

func (vs *VacancyService) GetAll(filters models.VacancyFilter) ([]*models.Vacancy, error) {
	var conditions []string
	var filterValues []interface{}
	values := reflect.ValueOf(filters)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface().(string) != "" {
			query := VacancyFilterQueries(types.Field(i).Name)
			if query != "" {
				conditions = append(conditions, query)
			}
			filterValues = append(filterValues, values.Field(i).Interface().(string))
		}
	}
	return vs.vacancyRep.GetAll(conditions, filterValues)
}

func (vs *VacancyService) Create(email string, input *models.Vacancy) (uint, error) {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return 0, getErr
	}

	if user.UserType != "employer" {
		return 0, errorHandler.InvalidUserType
	}
	input = escaping.EscapingObject[*models.Vacancy](input)

	input.PostedByUserId = user.ID
	return vs.vacancyRep.Create(input)
}

func (vs *VacancyService) GetById(vacancyID uint) (*models.Vacancy, error) {
	return vs.vacancyRep.GetById(vacancyID)
}

func (vs *VacancyService) GetByUserId(userId uint) ([]*models.Vacancy, error) {
	vacancies, getErr := vs.vacancyRep.GetByUserId(userId)
	if errors.Is(getErr, errorHandler.ErrVacancyNotFound) {
		return []*models.Vacancy{}, nil
	}
	return vacancies, getErr
}

func (vs *VacancyService) Delete(email string, vacancyId uint) error {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}
	userId := user.ID
	return vs.vacancyRep.Delete(userId, vacancyId)
}

func (vs *VacancyService) Update(email string, vacancyId uint, updates *models.Vacancy) error {

	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}
	userId := user.ID
	updates = escaping.EscapingObject[*models.Vacancy](updates)
	oldVacancy, getErr := vs.vacancyRep.GetById(vacancyId)
	if getErr != nil {
		return getErr
	}
	return vs.vacancyRep.Update(userId, vacancyId, oldVacancy, updates)
}
