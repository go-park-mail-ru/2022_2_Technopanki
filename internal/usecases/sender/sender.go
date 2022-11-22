package sender

import "HeadHunter/internal/entity/models"

type Sender interface {
	SendConfirmToken(email, token string) error
	SendApplicantMailing(email string, vacancies []*models.Vacancy) error
	SendEmployerMailing(email string, applicants []*models.UserAccount) error
}
