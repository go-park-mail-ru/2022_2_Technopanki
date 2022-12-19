package cron

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/usecases/mail"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
)

func FindApplicantsAndVacancies(db *gorm.DB) ([]string, []*models.VacancyPreview, error) {
	var applicants []*models.UserAccount
	query := db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "applicant", true, true).
		Scan(&applicants)
	if query.Error != nil {
		return nil, nil, query.Error
	}

	result := make([]string, len(applicants))
	for i, elem := range applicants {
		result[i] = elem.Email
	}

	var vacancies []*models.VacancyPreview
	query = db.Table("vacancies").Select("vacancies.title, vacancies.id, vacancies.location, user_accounts.company_name, vacancies.posted_by_user_id").
		Joins("left join user_accounts on user_accounts.id = vacancies.posted_by_user_id").
		Limit(5).Order("created_date desc").Scan(&vacancies)
	if query.Error != nil {
		return nil, nil, query.Error
	}

	return result, vacancies, nil
}

func FindEmployersAndResumes(db *gorm.DB) ([]string, []*models.ResumePreview, error) {
	var employers []*models.UserAccount
	query := db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "employer", true, true).
		Scan(&employers)
	if query.Error != nil {
		return nil, nil, query.Error
	}

	result := make([]string, len(employers))
	for i, elem := range employers {
		result[i] = elem.Email
	}

	var resumes []*models.ResumePreview
	query = db.Table("resumes").Select("resumes.title, resumes.id, resumes.location, user_accounts.applicant_name, resumes.user_account_id").
		Joins("left join user_accounts on user_accounts.id = resumes.user_account_id").
		Limit(5).Order("resumes.created_time desc").Scan(&resumes)
	if query.Error != nil {
		return nil, nil, query.Error
	}

	return result, resumes, nil
}

func Mailing(db *gorm.DB, mail mail.Mail) func() {
	return func() {
		employerEmails, applicants, employerErr := FindEmployersAndResumes(db)
		if employerErr == nil {
			err := mail.SendEmployerMailing(employerEmails, applicants)
			if err != nil {
				logrus.Println("send to applicant error: ", err)
			}
		}

		applicantEmails, vacancies, applicantErr := FindApplicantsAndVacancies(db)
		if applicantErr == nil {
			err := mail.SendApplicantMailing(applicantEmails, vacancies)
			if err != nil {
				logrus.Println("send to employer error: ", err)
			}
		}

		if employerErr != nil || applicantErr != nil {
			log.Println("errors with getting data: ", employerErr, applicantErr)
		}
	}
}
