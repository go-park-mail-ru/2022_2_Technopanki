package mail

import "HeadHunter/internal/entity/models"

type Mail interface {
	SendConfirmCode(email string) error
	SendApplicantMailing(users []*models.UserAccount, vacancies []*models.Vacancy) error
	SendEmployerMailing(employers, applicants []*models.UserAccount) error
}
