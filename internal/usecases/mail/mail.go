package mail

import "HeadHunter/internal/entity/models"

type Mail interface {
	SendConfirmCode(email string) error
	SendApplicantMailing(emails []string, vacancies []*models.VacancyPreview) error
	SendEmployerMailing(emails []string, previews []*models.ResumePreview) error
}
