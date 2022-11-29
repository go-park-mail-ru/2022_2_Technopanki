package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/handler"
	"HeadHunter/mail_microservice/usecase"
	"HeadHunter/pkg/errorHandler"
	"context"
)

type MailHandler struct {
	mailUseCase usecase.Mail
	handler.UnimplementedMailServiceServer
}

func NewMailHandler(_mail usecase.Mail) *MailHandler {
	return &MailHandler{mailUseCase: _mail}
}

func (mh *MailHandler) SendConfirmCode(ctx context.Context, in *handler.Email) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	sendErr := mh.mailUseCase.SendConfirmCode(in.Value)
	if sendErr != nil {
		return &handler.Nothing{}, sendErr
	}
	return &handler.Nothing{}, nil
}

func (mh *MailHandler) SendApplicantMailing(ctx context.Context, in *handler.ApplicantMailingData) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	vacancies := make([]*models.Vacancy, len(in.Vac))
	for i, preview := range in.Vac {
		vacancies[i].ID = uint(preview.Id)
		vacancies[i].Title = preview.Title
		vacancies[i].Image = preview.Image
	}
	sendErr := mh.mailUseCase.SendApplicantMailing(in.Emails, vacancies)
	if sendErr != nil {
		return &handler.Nothing{}, sendErr
	}
	return &handler.Nothing{}, nil
}

func (mh *MailHandler) SendEmployerMailing(ctx context.Context, in *handler.EmployerMailingData) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	applicants := make([]*models.UserAccount, len(in.Emp))
	for i, preview := range in.Emp {
		applicants[i].ID = uint(preview.Id)
		applicants[i].ApplicantName = preview.ApplicantName
		applicants[i].ApplicantSurname = preview.ApplicantSurname
		applicants[i].Image = preview.Image
		applicants[i].Status = preview.Status
		applicants[i].Location = preview.Location
	}
	sendErr := mh.mailUseCase.SendEmployerMailing(in.Emails, applicants)
	if sendErr != nil {
		return &handler.Nothing{}, sendErr
	}
	return &handler.Nothing{}, nil
}
