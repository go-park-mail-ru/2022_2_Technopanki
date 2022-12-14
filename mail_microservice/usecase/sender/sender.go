package sender

import "HeadHunter/internal/entity/models"

//go:generate mockgen -source=sender.go -destination=mocks/mock.go

type Sender interface {
	SendConfirmCode(email, code string) error
	SendApplicantMailing(email string, vacancies []models.VacancyPreview) error
	SendEmployerMailing(email string, previews []models.ResumePreview) error
}
