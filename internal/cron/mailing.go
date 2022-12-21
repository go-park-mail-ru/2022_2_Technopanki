package cron

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/usecases/mail"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
)

func findApplicantsToMailing(db *gorm.DB) ([]string, error) {
	var applicants []*models.UserAccount
	query := db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "applicant", true, true).
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

func findNewVacancies(db *gorm.DB) ([]*models.VacancyPreview, error) {
	var vacancies []*models.VacancyPreview
	query := db.Table("vacancies").Select("vacancies.title, vacancies.id, vacancies.location, user_accounts.company_name, vacancies.posted_by_user_id").
		Joins("left join user_accounts on user_accounts.id = vacancies.posted_by_user_id").
		Limit(5).Order("created_date desc").Scan(&vacancies)
	if query.Error != nil {
		return nil, query.Error
	}

	return vacancies, nil
}

func findEmployersToMailing(db *gorm.DB) ([]string, error) {
	var employers []*models.UserAccount
	query := db.Model(&models.UserAccount{}).Select("email").Where("user_type = ? AND mailing_approval = ? AND is_confirmed = ?", "employer", true, true).
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

func findNewResumes(db *gorm.DB) ([]*models.ResumePreview, error) {
	var resumes []*models.ResumePreview
	query := db.Table("resumes").Select("resumes.title, resumes.id, resumes.location, user_accounts.applicant_name, resumes.user_account_id").
		Joins("left join user_accounts on user_accounts.id = resumes.user_account_id").
		Limit(5).Order("resumes.created_time desc").Scan(&resumes)
	if query.Error != nil {
		return nil, query.Error
	}

	return resumes, nil
}

func Mailing(db *gorm.DB, mail mail.Mail) func() {
	return func() {
		employerEmails, employerErr := findEmployersToMailing(db)
		if employerErr != nil {
			log.Println("errors with getting employers emails: ", employerErr)
		}

		resumes, resumeErr := findNewResumes(db)
		if resumeErr != nil {
			log.Println("errors with getting resumes: ", resumeErr)
		}

		log.Println("emails: ", employerEmails, "\nnew resumes:")
		for _, e := range resumes {
			log.Println(*e)
		}

		err := mail.SendEmployerMailing(employerEmails, resumes)
		if err != nil {
			logrus.Println("send to applicant error: ", err)
		}

		applicantEmails, applicantErr := findApplicantsToMailing(db)
		if applicantErr != nil {
			log.Println("errors with getting applicants emails: ", applicantErr)
		}

		vacancies, vacancyErr := findNewVacancies(db)
		if vacancyErr != nil {
			log.Println("errors with getting vacancies: ", vacancyErr)
		}

		err = mail.SendApplicantMailing(applicantEmails, vacancies)
		if err != nil {
			logrus.Println("send to employer error: ", err)
		}
	}
}
