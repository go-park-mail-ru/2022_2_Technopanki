package usecase

import "HeadHunter/internal/entity/models"

type Mail interface {
	SendConfirmCode(email string) error
	SendApplicantMailing(emails []string, vacancies []*models.Vacancy) error
	SendEmployerMailing(emails []string, applicants []*models.UserAccount) error
}
