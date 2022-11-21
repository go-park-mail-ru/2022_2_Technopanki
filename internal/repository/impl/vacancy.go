package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"gorm.io/gorm"
)

type VacancyPostgres struct {
	db *gorm.DB
}

func NewVacancyPostgres(db *gorm.DB) *VacancyPostgres {
	return &VacancyPostgres{db: db}
}

func (vp *VacancyPostgres) GetAll() ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	query := vp.db.Find(&vacancies)
	if query.Error != nil {
		return vacancies, query.Error
	}
	return vacancies, nil
}

func (vp *VacancyPostgres) Create(vacancy *models.Vacancy) (uint, error) {
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

func (vp *VacancyPostgres) GetAuthor(email string) (*models.UserAccount, error) {
	return GetUser(email, vp.db)
}
