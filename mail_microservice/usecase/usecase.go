package usecase

import "HeadHunter/internal/entity/models"

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type Mail interface {
	SendConfirmCode(email string) error
	SendApplicantMailing(emails []string, vacancies []*models.VacancyPreview) error
	SendEmployerMailing(emails []string, previews []*models.ResumePreview) error
}
