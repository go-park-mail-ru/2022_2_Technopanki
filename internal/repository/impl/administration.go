package impl

import (
	"HeadHunter/internal/entity/models"
	"gorm.io/gorm"
)

type AdminPostgres struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (ap *AdminPostgres) GetResumesPageContent() ([]*models.ResumeAdminContent, error) {
	/*
		select resumes.id, resumes.title, max(user_accounts.applicant_surname), resumes.location, count(vacancy_activities)
		from resumes left join vacancy_activities
		on vacancy_activities.resume_id = resumes.id
		join user_accounts
		on resumes.user_account_id = user_accounts.id
		group by resumes.id
		order by resumes.id asc;
	*/
	var result []*models.ResumeAdminContent
	query := ap.db.Table("resumes").
		Select("resumes.id, resumes.title, max(user_accounts.applicant_surname) as applicant_surname, resumes.location, count(vacancy_activities) as responses").
		Joins("left join vacancy_activities on vacancy_activities.resume_id = resumes.id").
		Joins("join user_accounts on resumes.user_account_id = user_accounts.id").
		Group("resumes.id").Order("resumes.id asc").Scan(&result)
	return result, query.Error
}

func (ap *AdminPostgres) GetVacanciesPageContent() ([]*models.VacancyAdminContent, error) {
	/*
		select vacancies.id, vacancies.title, max(user_accounts.company_name), vacancies.location, count(vacancy_activities)
		from vacancies left join vacancy_activities
		on vacancy_activities.vacancy_id = vacancies.id
		join user_accounts
		on vacancies.posted_by_user_id = user_accounts.id
		group by vacancies.id
		order by vacancies.id asc;
	*/
	var result []*models.VacancyAdminContent
	query := ap.db.Table("vacancies").
		Select("vacancies.id, vacancies.title, max(user_accounts.company_name) as company_name, vacancies.location, count(vacancy_activities) as responses").
		Joins("left join vacancy_activities on vacancy_activities.vacancy_id = vacancies.id").
		Joins("join user_accounts on vacancies.posted_by_user_id = user_accounts.id").
		Group("vacancies.id").Order("vacancies.id asc").Scan(&result)
	return result, query.Error
}

func (ap *AdminPostgres) GetApplicantsPageContent() ([]*models.ApplicantAdminContent, error) {
	/*
		select user_accounts.id, user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.location, count(resumes)
		from user_accounts left join resumes
		on resumes.user_account_id = user_accounts.id
		where user_accounts.user_type = 'applicant'
		group by user_accounts.id
		order by user_accounts.id asc;
	*/
	var result []*models.ApplicantAdminContent
	query := ap.db.Table("user_accounts").
		Select("user_accounts.id, user_accounts.applicant_name, user_accounts.applicant_surname, user_accounts.location, count(resumes) as resumes_count").
		Joins("left join resumes on resumes.user_account_id = user_accounts.id").
		Where("user_accounts.user_type = 'applicant'").
		Group("user_accounts.id").Order("user_accounts.id asc").Scan(&result)
	return result, query.Error
}

func (ap *AdminPostgres) GetEmployersPageContent() ([]*models.EmployerAdminContent, error) {
	/*
		select user_accounts.id, user_accounts.company_name,  user_accounts.location, user_accounts.company_size, user_accounts.business_type, count(vacancies)
		from user_accounts left join vacancies
		on vacancies.posted_by_user_id = user_accounts.id
		where user_accounts.user_type = 'employer'
		group by user_accounts.id
		order by user_accounts.id asc;
	*/
	var result []*models.EmployerAdminContent
	query := ap.db.Table("user_accounts").
		Select("user_accounts.id, user_accounts.company_name,  user_accounts.location, user_accounts.company_size, user_accounts.business_type, count(vacancies) as vacancies_count").
		Joins("left join vacancies on vacancies.posted_by_user_id = user_accounts.id").
		Where("user_accounts.user_type = 'employer'").
		Group("user_accounts.id").Order("user_accounts.id asc").Scan(&result)
	return result, query.Error
}
