package sender

import "HeadHunter/internal/entity/models"

type Sender interface {
	SendConfirmCode(email, code string) error
	SendApplicantMailing(email string, vacancies []*models.Vacancy) error
	SendEmployerMailing(email string, applicants []*models.UserAccount) error
}
