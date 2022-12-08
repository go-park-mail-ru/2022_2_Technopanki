package impl

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
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

func (up *UserPostgres) UpdateUser(newUser *models.UserAccount) error {
	err := up.db.Model(newUser).Select("description", "two_factor_sign-in").Error
	if err != nil {
		return err
	}
	return up.db.Model(newUser).Updates(newUser).Error
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

func (up *UserPostgres) GetAllUsers(conditions []string, filterValues []interface{}, flag string) ([]*models.UserAccount, error) {
	var users []*models.UserAccount
	if conditions == nil {
		query := up.db.Where("user_type = ?", flag).Find(&users)
		if query.Error != nil {
			return users, query.Error
		}
		return users, nil

	} else {
		queryString := strings.Join(conditions, " AND ")
		queryConditions := FilterQueryStringFormatter(queryString, filterValues, up.db)
		query := queryConditions.Find(&users)
		if query.Error != nil {
			return users, query.Error
		}
		return users, nil
	}
}

func (up *UserPostgres) GetUserSafety(id uint, allowedFields []string) (*models.UserAccount, error) {
	var result models.UserAccount

	query := up.db.Select(append(models.SafeUserFields, allowedFields...)).Find(&result, id)
	return &result, QueryValidation(query, "user")
}

func (up *UserPostgres) GetBestVacanciesForApplicant() ([]*models.Vacancy, error) {
	var result []*models.Vacancy
	query := up.db.Table("vacancies").Order("created_date asc").Limit(10).Scan(&result)
	if query.Error != nil {
		return nil, fmt.Errorf("error with getting best vacancies: %w", query.Error)
	}
	return result, nil
}

func (up *UserPostgres) GetBestApplicantForEmployer() ([]*models.UserAccount, error) {
	var result []*models.UserAccount
	query := up.db.Table("user_accounts").Select(models.SafeUserFields).Where("user_accounts.user_type = ?", "applicant").
		Order("created_time asc").Limit(10).Scan(&result)
	if query.Error != nil {
		return nil, fmt.Errorf("error with getting best applicants: %w", query.Error)
	}
	return result, nil
}
