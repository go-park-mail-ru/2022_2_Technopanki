package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
	"strconv"
)

type VacancyPostgres struct {
	db *gorm.DB
}

func newVacancyPostgres(db *gorm.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}

func queryVacancyValidation(query *gorm.DB) error {
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errorHandler.ErrVacancyNotFound
	}
	return nil
}

func (vp *VacancyPostgres) GetAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	query := vp.db.Find(&vacancies)
	return vacancies, queryVacancyValidation(query)
}

func (vp *VacancyPostgres) Create(userId uint, vacancy *models.Vacancy) (uint, error) {
	vacancy.PostedByUserId = userId
	error := vp.db.Create(&vacancy).Error
	if error != nil {
		return 0, errorHandler.ErrInvalidQuery
	}
	return vacancy.ID, nil
}

func (vp *VacancyPostgres) GetById(vacancyId int) (*models.Vacancy, error) {
	var result models.Vacancy
	query := vp.db.Where("id = ?", vacancyId).Find(&result)
	return &result, queryVacancyValidation(query)

}

func (vp *VacancyPostgres) GetByUserId(userId uint) ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	query := vp.db.Where("PostedByUserId = ?", userId).Find(&vacancies)
	return vacancies, queryVacancyValidation(query)
}

func (vp *VacancyPostgres) Delete(userId uint, vacancyId int) error {

	error := vp.db.Where("PostedByUserId = ?", userId).Delete(&models.Vacancy{}, vacancyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancy
	}
	return nil
}

func (vp *VacancyPostgres) Update(userId uint, vacancyId int, vacancy *models.UpdateVacancy) error {
	userIdString := strconv.FormatUint(uint64(userId), 10)
	vacancyIdString := strconv.Itoa(vacancyId)

	error := vp.db.Model(&models.Vacancy{}).Where("ID = ? AND PostedByUserId = ?", vacancyIdString, userIdString).Updates(vacancy).Error
	if error != nil {
		return error
	}
	return error
}
