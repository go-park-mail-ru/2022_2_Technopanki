package mail

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/handler"
	"HeadHunter/pkg/errorHandler"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MailService struct {
	ctx    context.Context
	client handler.MailServiceClient
}

func NewMailService(_client handler.MailServiceClient) *MailService {
	return &MailService{
		client: _client,
		ctx:    context.Background(),
	}
}

func (ms *MailService) SendConfirmCode(email string) error {
	_, err := ms.client.SendConfirmCode(ms.ctx, &handler.Email{Value: email})
	st := status.Convert(err)
	if st.Code() == codes.AlreadyExists {
		return errorHandler.ErrCodeAlreadyExists
	}
	return st.Err()
}

func (ms *MailService) SendApplicantMailing(emails []string, vacancies []*models.VacancyPreview) error {
	if len(vacancies) < 5 || len(emails) == 0 {
		return nil
	}
	vacanciesPreview := make([]*handler.VacancyPreview, len(vacancies))
	for i, vacancy := range vacancies {
		vacanciesPreview[i] = &handler.VacancyPreview{}
		vacanciesPreview[i].Id = uint64(vacancy.Id)
		vacanciesPreview[i].Image = vacancy.Image
		vacanciesPreview[i].Title = vacancy.Title
	}

	_, err := ms.client.SendApplicantMailing(ms.ctx,
		&handler.ApplicantMailingData{Emails: emails, Vac: vacanciesPreview})
	if err != nil {
		return err
	}
	return nil
}

func (ms *MailService) SendEmployerMailing(emails []string, previews []*models.ResumePreview) error {
	if len(previews) < 5 || len(emails) == 0 {
		return nil
	}

	resumePreviews := make([]*handler.ResumePreview, len(previews))
	for i, preview := range previews {
		resumePreviews[i] = &handler.ResumePreview{}
		resumePreviews[i].ApplicantName = preview.ApplicantName
		resumePreviews[i].ApplicantSurname = preview.ApplicantSurname
		resumePreviews[i].UserAccountId = uint64(preview.UserAccountId)
		resumePreviews[i].Id = uint64(preview.Id)
		resumePreviews[i].Image = preview.Image
		resumePreviews[i].Title = preview.Title
		resumePreviews[i].Location = preview.Location
	}

	_, err := ms.client.SendEmployerMailing(ms.ctx, &handler.EmployerMailingData{Emails: emails, Emp: resumePreviews})
	if err != nil {
		return err
	}
	return nil
}
