package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
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

func (vp *VacancyPostgres) GetById(vacancyId int) (*models.Vacancy, error) {
	var result models.Vacancy
	query := vp.db.Where("id = ?", vacancyId).Find(&result)
	return &result, QueryValidation(query, "vacancy")
}

func (vp *VacancyPostgres) GetByUserId(userId int) ([]*models.Vacancy, error) {
	var vacancies []*models.Vacancy
	query := vp.db.Where("posted_by_user_id = ?", userId).Find(&vacancies)
	return vacancies, QueryValidation(query, "vacancy")
}

func (vp *VacancyPostgres) Delete(userId uint, vacancyId int) error {

	error := vp.db.Where("posted_by_user_id = ?", userId).Delete(&models.Vacancy{}, vacancyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancy
	}
	return nil
}

func (vp *VacancyPostgres) Update(userId uint, vacancyId int, oldVacancy *models.Vacancy, updates *models.Vacancy) error {

	query := vp.db.Model(oldVacancy).Where("id = ? AND posted_by_user_id = ?", vacancyId, userId).Updates(updates)
	return QueryValidation(query, "vacancy")
}

//func (vp *VacancyPostgres) Update(vacancyId int, updates *models.Vacancy) error {
//old, getErr := vp.GetById(vacancyId)
//fmt.Println(old)
//if getErr != nil {
//return getErr
//}
//updates.PostedByUserId = old.PostedByUserId
//updates.ID = uint(vacancyId)
//return vp.db.Updates(updates).Error
//}
