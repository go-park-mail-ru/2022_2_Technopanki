package mail

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/handler"
	"HeadHunter/pkg/errorHandler"
	"context"
	"fmt"
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
	fmt.Println(st.Code())
	fmt.Println(st.Code() == codes.AlreadyExists)
	fe, _ := status.FromError(err)
	fmt.Println(fe.Code(), fe.String())
	if err != nil && err.Error() == "rpc error: code = Unknown desc = Код уже отправлен" { //TODO переделать
		return errorHandler.ErrCodeAlreadyExists
	}
	return err
}

func (ms *MailService) SendApplicantMailing(emails []string, vacancies []*models.Vacancy) error {
	vacanciesPreview := make([]*handler.VacancyPreview, len(vacancies))
	for i, vacancy := range vacancies {
		vacanciesPreview[i].Id = uint64(vacancy.ID)
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

func (ms *MailService) SendEmployerMailing(emails []string, applicants []*models.UserAccount) error {
	applicantPreview := make([]*handler.ApplicantPreview, len(applicants))
	for i, applicant := range applicants {
		applicantPreview[i].ApplicantName = applicant.ApplicantName
		applicantPreview[i].ApplicantSurname = applicant.ApplicantSurname
		applicantPreview[i].Id = uint64(applicant.ID)
		applicantPreview[i].Image = applicant.Image
		applicantPreview[i].Status = applicant.Status
		applicantPreview[i].Location = applicant.Location
	}
	_, err := ms.client.SendEmployerMailing(ms.ctx, &handler.EmployerMailingData{Emails: emails})
	if err != nil {
		return err
	}
	return nil
}
