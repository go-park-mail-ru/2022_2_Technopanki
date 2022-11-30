package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"gorm.io/gorm"
	"strings"
)

type VacancyPostgres struct {
	db *gorm.DB
}

func NewVacancyPostgres(db *gorm.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}

func (vp *VacancyPostgres) GetAll(conditions []string, filterValues []interface{}) ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	if conditions == nil {
		query := vp.db.Find(&vacancies)
		if query.Error != nil {
			return vacancies, query.Error
		}
		return vacancies, nil

	} else {
		queryString := strings.Join(conditions, " AND ")
		queryConditions := FilterQueryStringFormatter(queryString, filterValues, vp.db)
		query := queryConditions.Find(&vacancies)
		if query.Error != nil {
			return vacancies, query.Error
		}
		return vacancies, nil
	}
}

func (vp *VacancyPostgres) GetAllFilter(filter string) ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	query := vp.db.Where("title LIKE ?", filter).Find(&vacancies)
	if query.Error != nil {
		return vacancies, query.Error
	}
	return vacancies, nil
}

func (vp *VacancyPostgres) Create(vacancy *models.Vacancy) (uint, error) {
	var user models.UserAccount
	queryUser := vp.db.Where("id = ?", vacancy.PostedByUserId).Find(&user)
	if queryUser.Error != nil {
		return 0, queryUser.Error
	}
	vacancy.Image = user.Image
	error := vp.db.Create(vacancy).Error
	if error != nil {
		return 0, errorHandler.ErrInvalidParam
	}
	return vacancy.ID, nil
}

func (vp *VacancyPostgres) GetById(vacancyId uint) (*models.Vacancy, error) {
	var result models.Vacancy
	query := vp.db.Where("id = ?", vacancyId).Find(&result)
	return &result, QueryValidation(query, "vacancy")
}

func (vp *VacancyPostgres) GetByUserId(userId uint) ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	query := vp.db.Where("posted_by_user_id = ?", userId).Find(&vacancies)
	return vacancies, QueryValidation(query, "vacancy")
}

func (vp *VacancyPostgres) Delete(userId uint, vacancyId uint) error {

	error := vp.db.Where("posted_by_user_id = ?", userId).Delete(&models.Vacancy{}, vacancyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancy
	}
	return nil
}

func (vp *VacancyPostgres) Update(userId uint, vacancyId uint, oldVacancy *models.Vacancy, updates *models.Vacancy) error {

	query := vp.db.Model(oldVacancy).Where("id = ? AND posted_by_user_id = ?", vacancyId, userId).Updates(updates)
	return QueryValidation(query, "vacancy")
}
