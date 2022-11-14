package repository

import (
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
)

type VacancyActivityPostgres struct {
	db *gorm.DB
}

func newVacancyActivityPostgres(db *gorm.DB) *VacancyActivityPostgres {
	return &VacancyActivityPostgres{db: db}
}

func (vap *VacancyActivityPostgres) GetAllVacancyApplies(vacancyId int) ([]*models.VacancyActivity, error) {
	var applies []*models.VacancyActivity
	//var responce []*models.VacancyActivityResponce
	//query := vap.db.Model(&models.Resume{}).Select("resumes.title, resumes.description").Joins("left join vacancy_activities on vacancy_activities.resume_id = resumes.id").Where("vacancy_id = ?", vacancyId).Scan(&responce)
	query := vap.db.Where("vacancy_id = ?", vacancyId).Find(&applies)
	if query.Error != nil {
		return applies, query.Error
	}
	return applies, nil
}

func (vap *VacancyActivityPostgres) ApplyForVacancy(apply *models.VacancyActivity) error {
	var user models.UserAccount
	queryUser := vap.db.Where("id = ?", apply.UserAccountId).Find(&user)
	if queryUser.Error != nil {
		return queryUser.Error
	}
	apply.Image = user.Image
	query := vap.db.Create(&apply)
	//return queryValidation(query, "vacancy_activity")
	if query.RowsAffected == 0 {
		return notFound("vacancy_activity")
	}
	return queryValidation(query, "vacancy_activity")
}

func (vap *VacancyActivityPostgres) GetAllUserApplies(userId int) ([]*models.VacancyActivity, error) {
	var applies []*models.VacancyActivity
	query := vap.db.Where("user_account_id = ?", userId).Find(&applies)
	if query.Error != nil {
		return applies, query.Error
	}
	return applies, nil
}
