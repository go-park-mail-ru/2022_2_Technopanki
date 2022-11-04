package repository

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/entity/validation"
	"HeadHunter/internal/errorHandler"
	"gorm.io/gorm"
)

type VacancyActivityPostgres struct {
	db *gorm.DB
}

func newVacancyActivityPostgres(db *gorm.DB) *VacancyActivityPostgres {
	return &VacancyActivityPostgres{db: db}
}

func (vap *VacancyActivityPostgres) GetAllVacancyApplies(vacancyId int) ([]models.VacancyActivity, error) {
	var applies []models.VacancyActivity
	query := vap.db.Where("vacancy_id = ?", vacancyId).Find(applies)
	return applies, validation.QueryVacancyValidation(query)
}

func (vap *VacancyActivityPostgres) ApplyForVacancy(userId uint, apply *models.VacancyActivity) error {
	apply.UserAccountId = userId
	error := vap.db.Create(&apply).Error
	if error != nil {
		return errorHandler.ErrInvalidParam
	}
	return nil
}
