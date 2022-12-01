package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/pkg/errorHandler"
	"gorm.io/gorm"
)

type VacancyActivityPostgres struct {
	db *gorm.DB
}

func NewVacancyActivityPostgres(db *gorm.DB) *VacancyActivityPostgres {
	return &VacancyActivityPostgres{db: db}
}

func (vap *VacancyActivityPostgres) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error) { //TODO Переделать
	var applies []*models.VacancyActivity
	query := vap.db.Where("vacancy_id = ?", vacancyId).Find(&applies)
	if query.Error != nil {
		return applies, query.Error
	}
	var user models.UserAccount
	if len(applies) > 0 {
		queryUser := vap.db.Where("id = ?", applies[0].UserAccountId).Find(&user)
		if queryUser.Error != nil {
			return applies, queryUser.Error
		}
	}
	for _, elem := range applies {
		elem.Image = user.Image
		elem.ApplicantName = user.ApplicantName
		elem.ApplicantSurname = user.ApplicantSurname
	}
	return applies, nil
}

func (vap *VacancyActivityPostgres) ApplyForVacancy(apply *models.VacancyActivity) error { //TODO переделать
	var user models.UserAccount
	queryUser := vap.db.Where("id = ?", apply.UserAccountId).Find(&user)
	if queryUser.Error != nil {
		return queryUser.Error
	}
	var resume models.Vacancy
	queryVacancy := vap.db.Where("id = ?", apply.ResumeId).Find(&resume)
	if queryVacancy.Error != nil {
		return queryVacancy.Error
	}

	apply.ResumeTitle = resume.Title
	apply.Image = user.Image
	query := vap.db.Create(&apply)
	return QueryValidation(query, "vacancy_activity")
}

func (vap *VacancyActivityPostgres) GetAllUserApplies(userId uint) ([]*models.VacancyActivity, error) { //TODO переделать
	var applies []*models.VacancyActivity
	var user models.UserAccount
	queryUser := vap.db.Where("id = ?", userId).Find(&user)
	if queryUser.Error != nil {
		return applies, queryUser.Error
	}

	query := vap.db.Where("user_account_id = ?", userId).Find(&applies)
	if query.Error != nil {
		return applies, query.Error
	}

	for _, elem := range applies {
		elem.Image = user.Image
		elem.ApplicantName = user.ApplicantName
		elem.ApplicantSurname = user.ApplicantSurname
	}
	return applies, nil
}

func (vap *VacancyActivityPostgres) DeleteUserApply(userId uint, applyId uint) error {
	error := vap.db.Where("user_account_id = ?", userId).Delete(&models.VacancyActivity{}, applyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancyApply
	}
	return nil
}
