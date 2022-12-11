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

func (vp *VacancyPostgres) GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error) {
	var resultPreview []*models.VacancyPreview
	query := vp.db.Table("vacancies").
		Select("user_accounts.image,"+
			"vacancies.id, vacancies.title, vacancies.salary, vacancies.location, vacancies.format, vacancies.hours, vacancies.description").
		Joins("left join user_accounts on vacancies.posted_by_user_id = user_accounts.id").
		Where("posted_by_user_id = ?", userId).
		Scan(&resultPreview)
	return resultPreview, QueryValidation(query, "vacancy")

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

func (vp *VacancyPostgres) AddVacancyToFavorites(user *models.UserAccount, vacancy *models.Vacancy) error {
	err := vp.db.Model(user).Association("FavoriteVacancies").Append(vacancy)
	if err != nil {
		return errorHandler.ErrCannotAddFavoriteVacancy
	}
	return nil
}

func (vp *VacancyPostgres) GetUserFavoriteVacancies(user *models.UserAccount) ([]*models.Vacancy, error) {
	var favoriteVacancies []*models.Vacancy
	err := vp.db.Model(user).Association("FavoriteVacancies").Find(&favoriteVacancies)
	if err != nil {
		return nil, errorHandler.ErrCannotGetFavoriteVacancy
	}
	return favoriteVacancies, nil
}

func (vp *VacancyPostgres) DeleteVacancyFromFavorites(user *models.UserAccount, vacancy *models.Vacancy) error {
	err := vp.db.Model(user).Association("FavoriteVacancies").Delete(vacancy)
	if err != nil {
		return errorHandler.ErrCannotDeleteVacancyFromFavorites
	}
	return nil
}
