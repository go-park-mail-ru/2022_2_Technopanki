package impl

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateUser(user *models.UserAccount) error {
	return up.db.Create(user).Error
}

func (up *UserPostgres) UpdateUser(oldUser, newUser *models.UserAccount) error {
	return up.db.Model(oldUser).Updates(newUser).Error
}
func (up *UserPostgres) UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error {
	return up.db.Model(oldUser).Select(field).Updates(newUser).Error
}
func (up *UserPostgres) GetUserByEmail(email string) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Where("email = ?", email).Find(&result)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) GetUser(id uint) (*models.UserAccount, error) {
	var result models.UserAccount
	query := up.db.Select(append(models.PrivateUserFields, models.SafeUserFields...)).Find(&result, id)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) GetUserSafety(id uint, allowedFields []string) (*models.UserAccount, error) {
	var result models.UserAccount

	query := up.db.Select(append(models.SafeUserFields, allowedFields...)).Find(&result, id)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) UpdatePassword(user *models.UserAccount) error {
	query := up.db.Model(user).Select("password").Updates(user)
	return query.Error
}

func (up *UserPostgres) GetBestVacanciesForApplicant(user *models.UserAccount) ([]*models.Vacancy, error) {
	var result []*models.Vacancy
	query := up.db.Table("vacancies").Order("created_date asc").Limit(10).Scan(&result)
	if query.Error != nil {
		return result, fmt.Errorf("error with getting best vacancies: %w", query.Error)
	}
	return result, nil
}

func (up *UserPostgres) GetBestApplicantForEmployer(user *models.UserAccount) ([]*models.UserAccount, error) {
	var result []*models.UserAccount
	query := up.db.Table("user_accounts").Select(models.SafeUserFields).Where("user_accounts.user_type = ?", "applicant").
		Order("created_time desc").Limit(10).Scan(&result)
	if query.Error != nil {
		return result, fmt.Errorf("error with getting best applicants: %w", query.Error)
	}
	return result, nil
}
