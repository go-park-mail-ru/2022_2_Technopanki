package sender

import "HeadHunter/internal/entity/models"

type Sender interface {
	SendConfirmToken(token string) error
	SendApplicantMailing(email string, vacancies []*models.Vacancy) error
	SendEmployerMailing(email string, applicants []*models.UserAccount) error
}
