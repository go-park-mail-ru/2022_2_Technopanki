package impl

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type ResumePostgres struct {
	db *gorm.DB
}

func NewResumePostgres(db *gorm.DB) *ResumePostgres {
	return &ResumePostgres{db: db}
}

func (rp *ResumePostgres) GetResume(id uint) (*models.Resume, error) {
	var result models.Resume

	query := rp.db.Table("resumes").
		Joins("left join experience_details on resumes.id = experience_details.resume_id").
		Joins("left join education_details on resumes.id = education_details.resume_id").
		Where("resumes.id = ?", id).Scan(&result).
		Scan(&result.ExperienceDetail).Scan(&result.EducationDetail)

	return &result, QueryValidation(query, "resume")
}

func (rp *ResumePostgres) GetAllResumes(conditions []string, filterValues []interface{}) ([]*models.Resume, error) {
	var resumes []*models.Resume
	if conditions == nil {
		query := rp.db.Find(&resumes)
		if query.Error != nil {
			return resumes, query.Error
		}
		return resumes, nil

	} else {
		queryString := strings.Join(conditions, " AND ")
		queryConditions := FilterQueryStringFormatter(queryString, filterValues, rp.db)
		query := queryConditions.Find(&resumes)
		if query.Error != nil {
			return resumes, query.Error
		}
		return resumes, nil
	}
}

func (rp *ResumePostgres) GetResumeByApplicant(userId uint) ([]*models.Resume, error) {
	var result []*models.Resume
	var resultEdu []*models.EducationDetail
	var resultExp []*models.ExperienceDetail
	query := rp.db.Table("resumes").
		Joins("left join experience_details on resumes.id = experience_details.resume_id").
		Joins("left join education_details on resumes.id = education_details.resume_id").
		Where("user_account_id = ?", userId).
		Scan(&result).Scan(&resultEdu).Scan(&resultExp)

	for i, elem := range result {
		elem.ExperienceDetail = *resultExp[i]
		resultExp[i].ResumeId = elem.ID

		elem.EducationDetail = *resultEdu[i]
		resultEdu[i].ResumeId = elem.ID
	}
	return result, QueryValidation(query, "resume")
}

func (rp *ResumePostgres) GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error) {
	var resultPreview []*models.ResumePreview
	query := rp.db.Table("resumes").
		Select("user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.image,"+
			"resumes.id, resumes.title, resumes.created_time").
		Joins("left join user_accounts on resumes.user_account_id = user_accounts.id").
		Where("user_account_id = ?", userId).
		Scan(&resultPreview)

	return resultPreview, QueryValidation(query, "resume")
}

func (rp *ResumePostgres) GetResumeInPDF(resumeId uint) (*models.ResumeInPDF, error) {
	var result *models.ResumeInPDF
	query := rp.db.Table("resumes").
		Select("user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.image,"+
			"user_accounts.contact_number, user_accounts.email, user_accounts.location, resumes.title, resumes.description").
		Joins("left join user_accounts on resumes.user_account_id = user_accounts.id").
		Where("resumes.id = ?", resumeId).
		Scan(&result)

	return result, QueryValidation(query, "resume")
}

func (rp *ResumePostgres) CreateResume(resume *models.Resume, userId uint) error {
	resume.UserAccountId = userId

	var user models.UserAccount
	queryUser := rp.db.Where("id = ?", userId).Find(&user)

	if queryUser.Error != nil {
		return queryUser.Error
	}

	resume.Location = user.Location

	creatingErr := rp.db.Create(resume).Save(resume).Error

	if creatingErr != nil {
		return fmt.Errorf("cannot create resume: %w", creatingErr)
	}

	return nil
}

func (rp *ResumePostgres) UpdateResume(id uint, resume *models.Resume) error {
	resume.ID = id
	return rp.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(resume).Error
}

func (rp *ResumePostgres) DeleteResume(id uint) error {
	return rp.db.Delete(&models.Resume{ID: id}).Error
}

func (rp *ResumePostgres) GetEmployerIdByVacancyActivity(id uint) (uint, error) {
	var result uint
	query := rp.db.Model(&models.Resume{}).Select("vacancies.posted_by_user_id").
		Joins("left join vacancy_activities on resumes.id = vacancy_activities.resume_id").
		Joins("left join vacancies on vacancies.id = vacancy_activities.vacancy_id").
		Where("resumes.id = ?", id).
		Scan(&result)
	return result, QueryValidation(query, "user")
}
