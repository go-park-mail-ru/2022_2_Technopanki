package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/pkg/errorHandler"
	"bytes"
	"html/template"
	"io/ioutil"
)

type AdministrationService struct {
	adminRep repository.Administration
	userRep  repository.UserRepository
}

func NewAdministrationService(adminRepository repository.Administration, userRepository repository.UserRepository) *AdministrationService {
	return &AdministrationService{adminRep: adminRepository, userRep: userRepository}
}

func (as *AdministrationService) GetMainPage(email string) ([]byte, error) {
	user, getErr := as.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if !user.IsAdmin {
		return nil, errorHandler.ErrForbidden
	}
	file, readErr := ioutil.ReadFile("static/html/admin/main.html")
	if readErr != nil {
		return nil, readErr
	}
	return file, nil
}
func (as *AdministrationService) GetAuthPage() ([]byte, error) {
	file, readErr := ioutil.ReadFile("/usr/share/html/admin/auth.html")
	if readErr != nil {
		return nil, readErr
	}
	return file, nil
}
func (as *AdministrationService) GetResumesPage(email string) ([]byte, error) {
	user, getErr := as.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if !user.IsAdmin {
		return nil, errorHandler.ErrForbidden
	}

	resumes, resumeErr := as.adminRep.GetResumesPageContent()
	if resumeErr != nil {
		return nil, resumeErr
	}
	resumesValid := make([]models.ResumeAdminContent, len(resumes), len(resumes))
	for i, resume := range resumes {
		resumesValid[i] = *resume
	}
	data := struct {
		Resumes []models.ResumeAdminContent
	}{
		Resumes: resumesValid,
	}

	form, parseErr := template.ParseFiles("static/html/admin/resumes.html")
	if parseErr != nil {
		return nil, parseErr
	}

	formBuf := bytes.NewBuffer([]byte(""))
	executeErr := form.Execute(formBuf, data)
	if executeErr != nil {
		return nil, executeErr
	}
	return formBuf.Bytes(), nil

}
func (as *AdministrationService) GetVacanciesPage(email string) ([]byte, error) {
	user, getErr := as.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if !user.IsAdmin {
		return nil, errorHandler.ErrForbidden
	}

	vacancies, vacancyErr := as.adminRep.GetVacanciesPageContent()
	if vacancyErr != nil {
		return nil, vacancyErr
	}
	vacanciesValid := make([]models.VacancyAdminContent, len(vacancies), len(vacancies))
	for i, vacancy := range vacancies {
		vacanciesValid[i] = *vacancy
	}
	data := struct {
		Vacancies []models.VacancyAdminContent
	}{
		Vacancies: vacanciesValid,
	}

	form, parseErr := template.ParseFiles("static/html/admin/vacancies.html")
	if parseErr != nil {
		return nil, parseErr
	}

	formBuf := bytes.NewBuffer([]byte(""))
	executeErr := form.Execute(formBuf, data)
	if executeErr != nil {
		return nil, executeErr
	}
	return formBuf.Bytes(), nil
}
func (as *AdministrationService) GetApplicantsPage(email string) ([]byte, error) {
	user, getErr := as.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if !user.IsAdmin {
		return nil, errorHandler.ErrForbidden
	}

	applicants, applicantErr := as.adminRep.GetApplicantsPageContent()
	if applicantErr != nil {
		return nil, applicantErr
	}
	applicantsValid := make([]models.ApplicantAdminContent, len(applicants), len(applicants))
	for i, applicant := range applicants {
		applicantsValid[i] = *applicant
	}
	data := struct {
		Applicants []models.ApplicantAdminContent
	}{
		Applicants: applicantsValid,
	}

	form, parseErr := template.ParseFiles("static/html/admin/applicants.html")
	if parseErr != nil {
		return nil, parseErr
	}

	formBuf := bytes.NewBuffer([]byte(""))
	executeErr := form.Execute(formBuf, data)
	if executeErr != nil {
		return nil, executeErr
	}
	return formBuf.Bytes(), nil
}
func (as *AdministrationService) GetEmployersPage(email string) ([]byte, error) {
	user, getErr := as.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if !user.IsAdmin {
		return nil, errorHandler.ErrForbidden
	}
	employers, employerErr := as.adminRep.GetEmployersPageContent()
	if employerErr != nil {
		return nil, employerErr
	}
	employersValid := make([]models.EmployerAdminContent, len(employers), len(employers))
	for i, employer := range employers {
		employersValid[i] = *employer
	}
	data := struct {
		Employers []models.EmployerAdminContent
	}{
		Employers: employersValid,
	}

	form, parseErr := template.ParseFiles("static/html/admin/employers.html")
	if parseErr != nil {
		return nil, parseErr
	}

	formBuf := bytes.NewBuffer([]byte(""))
	executeErr := form.Execute(formBuf, data)
	if executeErr != nil {
		return nil, executeErr
	}
	return formBuf.Bytes(), nil
}
