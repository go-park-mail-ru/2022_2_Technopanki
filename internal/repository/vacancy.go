package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
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

func (vp *VacancyPostgres) GetAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	query := vp.db.Find(&vacancies)
	return vacancies, validation.QueryVacancyValidation(query)
}

func (vp *VacancyPostgres) Create(userId uint, vacancy *models.Vacancy) (uint, error) {
	vacancy.PostedByUserId = userId
	error := vp.db.Create(&vacancy).Error
	if error != nil {
		return 0, errorHandler.ErrInvalidParam
	}
	return vacancy.ID, nil
}

func (vp *VacancyPostgres) GetById(vacancyId int) (*models.Vacancy, error) {
	var result models.Vacancy
	query := vp.db.Where("id = ?", vacancyId).Find(&result)
	return &result, validation.QueryVacancyValidation(query)

}

func (vp *VacancyPostgres) GetByUserId(userId int) ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	query := vp.db.Where("posted_by_user_id = ?", userId).Find(&vacancies)
	return vacancies, validation.QueryVacancyValidation(query)
}

func (vp *VacancyPostgres) Delete(userId uint, vacancyId int) error {

	error := vp.db.Where("posted_by_user_id = ?", userId).Delete(&models.Vacancy{}, vacancyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancy
	}
	return nil
}

func (vp *VacancyPostgres) Update(userId uint, vacancyId int, vacancy *models.UpdateVacancy) error {
	userIdString := strconv.FormatUint(uint64(userId), 10)
	vacancyIdString := strconv.Itoa(vacancyId)

	error := vp.db.Model(&models.Vacancy{}).Where("id = ? AND posted_by_user_id = ?", vacancyIdString, userIdString).Updates(vacancy).Error
	if error != nil {
		return error
	}
	return error
}
