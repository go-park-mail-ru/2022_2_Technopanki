package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/usecases/escaping"
	"HeadHunter/pkg/errorHandler"
	"errors"
	"reflect"
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

func (vs *VacancyService) GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error) {
	vacanciesPreview, getErr := vs.vacancyRep.GetPreviewVacanciesByEmployer(userId)
	if errors.Is(getErr, errorHandler.ErrVacancyNotFound) {
		return []*models.VacancyPreview{}, nil
	}
	return vacanciesPreview, nil
}

func (vs *VacancyService) GetByUserId(userId uint) ([]*models.Vacancy, error) {
	return vs.vacancyRep.GetByUserId(userId)
}

func (vs *VacancyService) Delete(email string, vacancyId uint) error {
	user, getUserErr := vs.userRep.GetUserByEmail(email)
	if getUserErr != nil {
		return getUserErr
	}

	old, getErr := vs.vacancyRep.GetById(vacancyId)
	if getErr != nil {
		return getErr
	}

	if user.ID != old.PostedByUserId && !user.IsAdmin {
		return errorHandler.ErrUnauthorized
	}

	return vs.vacancyRep.Delete(user.ID, vacancyId)
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
func (vs *VacancyService) AddVacancyToFavorites(email string, vacancyId uint) error {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}

	vacancy, vacancyErr := vs.vacancyRep.GetById(vacancyId)
	if vacancyErr != nil {
		return vacancyErr
	}
	return vs.vacancyRep.AddVacancyToFavorites(user, vacancy)
}

func (vs *VacancyService) GetUserFavoriteVacancies(email string) ([]*models.Vacancy, error) {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	return vs.vacancyRep.GetUserFavoriteVacancies(user)
}

func (vs *VacancyService) DeleteVacancyFromFavorites(email string, vacancyId uint) error {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}
	vacancy, err := vs.vacancyRep.GetById(vacancyId)
	if err != nil {
		return err
	}
	return vs.vacancyRep.DeleteVacancyFromFavorites(user, vacancy)

}

func (vs *VacancyService) CheckFavoriteVacancy(email string, vacancyId uint) (bool, error) {
	user, getErr := vs.userRep.GetUserByEmail(email)
	if getErr != nil {
		return false, getErr
	}
	userId := user.ID
	return vs.vacancyRep.CheckFavoriteVacancy(userId, vacancyId)
}
