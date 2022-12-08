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

func (vap *VacancyActivityPostgres) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivityPreview, error) {
	var appliesPreview []*models.VacancyActivityPreview
	query := vap.db.Table("vacancy_activities").
		Select("user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.image,"+
			"vacancy_activities.user_account_id, vacancy_activities.resume_id, vacancy_activities.vacancy_id, vacancy_activities.apply_date,"+
			"resumes.title").
		Joins("left join user_accounts on user_account.id = vacancy_activities.user_account_id").
		Joins("left join resumes on resumes.id = vacancy_activities.resume_id").
		Where("vacancy_id = ?", vacancyId).
		Scan(&appliesPreview)

	return appliesPreview, QueryValidation(query, "vacancy_applies")
}

func (vap *VacancyActivityPostgres) ApplyForVacancy(apply *models.VacancyActivity) error {
	query := vap.db.Create(&apply)
	return QueryValidation(query, "vacancy_activity")
}

func (vap *VacancyActivityPostgres) GetAllUserApplies(userId uint) ([]*models.VacancyActivityPreview, error) {
	var appliesPreview []*models.VacancyActivityPreview
	query := vap.db.Table("vacancy_activities").
		Select("user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.image,"+
			"vacancy_activities.user_account_id, vacancy_activities.resume_id, vacancy_activities.vacancy_id, vacancy_activities.apply_date,"+
			"resumes.title").
		Joins("left join user_accounts on user_account.id = vacancy_activities.user_account_id").
		Joins("left join resumes on resumes.id = vacancy_activities.resume_id").
		Where("user_account_id = ?", userId).
		Scan(&appliesPreview)

	return appliesPreview, QueryValidation(query, "vacancy_applies")
}

func (vap *VacancyActivityPostgres) DeleteUserApply(userId uint, applyId uint) error {
	error := vap.db.Where("user_account_id = ?", userId).Delete(&models.VacancyActivity{}, applyId).Error
	if error != nil {
		return errorHandler.ErrCannotDeleteVacancyApply
	}
	return nil
}
