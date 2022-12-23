package impl

import (
	"HeadHunter/internal/entity/models"
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
	err := up.db.Model(newUser).Updates(map[string]interface{}{"two_factor_sign_in": newUser.TwoFactorSignIn,
		"description": newUser.Description, "mailing_approval": newUser.MailingApproval}).Error
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
		query := up.db.Select(models.SafeUserFields).Where("user_type = ?", flag).Find(&users)
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

func (up *UserPostgres) FindApplicantsToMailing() ([]string, error) {
	var applicants []*models.UserAccount
	query := up.db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "applicant", true, true).
		Scan(&applicants)
	if query.Error != nil {
		return nil, query.Error
	}

	result := make([]string, len(applicants))
	for i, elem := range applicants {
		result[i] = elem.Email
	}
	return result, nil
}

func (up *UserPostgres) FindNewVacancies() ([]*models.VacancyPreview, error) {
	var vacancies []*models.VacancyPreview
	query := up.db.Table("vacancies").Select("vacancies.title, vacancies.id, vacancies.location, user_accounts.company_name, vacancies.posted_by_user_id").
		Joins("left join user_accounts on user_accounts.id = vacancies.posted_by_user_id").
		Limit(5).Order("created_date desc").Scan(&vacancies)
	if query.Error != nil {
		return nil, query.Error
	}

	return vacancies, nil
}

func (up *UserPostgres) FindEmployersToMailing() ([]string, error) {
	var employers []*models.UserAccount
	query := up.db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "employer", true, true).
		Scan(&employers)
	if query.Error != nil {
		return nil, query.Error
	}

	result := make([]string, len(employers))
	for i, elem := range employers {
		result[i] = elem.Email
	}
	return result, nil
}

func (up *UserPostgres) FindNewResumes() ([]*models.ResumePreview, error) {
	var resumes []*models.ResumePreview
	query := up.db.Table("resumes").Select("resumes.title, resumes.id, resumes.location, user_accounts.applicant_name, resumes.user_account_id").
		Joins("left join user_accounts on user_accounts.id = resumes.user_account_id").
		Limit(5).Order("resumes.created_time desc").Scan(&resumes)
	if query.Error != nil {
		return nil, query.Error
	}

	return resumes, nil
}
