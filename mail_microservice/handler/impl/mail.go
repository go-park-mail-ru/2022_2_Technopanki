package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/handler"
	"HeadHunter/mail_microservice/usecase"
	"HeadHunter/metrics"
	"HeadHunter/pkg/errorHandler"
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type MailHandler struct {
	mailUseCase usecase.Mail
	handler.UnimplementedMailServiceServer
}

func NewMailHandler(_mail usecase.Mail) *MailHandler {
	return &MailHandler{mailUseCase: _mail}
}

func (mh *MailHandler) SendConfirmCode(ctx context.Context, in *handler.Email) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.MailRequestDuration.WithLabelValues("SendConfirmCode"))
	defer timer.ObserveDuration()
	if in == nil {
		metrics.MailRequest.WithLabelValues("400", "bad request", "SendConfirmCode").Inc()
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	sendErr := mh.mailUseCase.SendConfirmCode(in.Value)
	if sendErr != nil {
		metrics.MailRequest.WithLabelValues("500", "send message error", "SendConfirmCode").Inc()
		return &handler.Nothing{}, sendErr
	}
	metrics.MailRequest.WithLabelValues("200", "success", "SendConfirmCode").Inc()
	return &handler.Nothing{}, nil
}

func (mh *MailHandler) SendApplicantMailing(ctx context.Context, in *handler.ApplicantMailingData) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.MailRequestDuration.WithLabelValues("SendApplicantMailing"))
	defer timer.ObserveDuration()
	if in == nil {
		metrics.MailRequest.WithLabelValues("400", "bad request", "SendApplicantMailing").Inc()
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	vacancies := make([]*models.VacancyPreview, len(in.Vac))
	for i, preview := range in.Vac {
		vacancies[i] = &models.VacancyPreview{}
		vacancies[i].Id = uint(preview.Id)
		vacancies[i].Title = preview.Title
		vacancies[i].Image = preview.Image
	}
	sendErr := mh.mailUseCase.SendApplicantMailing(in.Emails, vacancies)
	if sendErr != nil {
		metrics.MailRequest.WithLabelValues("500", "send message error", "SendApplicantMailing").Inc()
		return &handler.Nothing{}, sendErr
	}
	metrics.MailRequest.WithLabelValues("200", "success", "SendApplicantMailing").Inc()
	return &handler.Nothing{}, nil
}

func (mh *MailHandler) SendEmployerMailing(ctx context.Context, in *handler.EmployerMailingData) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.MailRequestDuration.WithLabelValues("SendEmployerMailing"))
	defer timer.ObserveDuration()
	if in == nil {
		metrics.MailRequest.WithLabelValues("400", "bad request", "SendEmployerMailing").Inc()
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	previews := make([]*models.ResumePreview, len(in.Emp))
	for i, preview := range in.Emp {
		previews[i] = &models.ResumePreview{}
		previews[i].Id = uint(preview.Id)
		previews[i].UserAccountId = uint(preview.UserAccountId)
		previews[i].ApplicantName = preview.ApplicantName
		previews[i].ApplicantSurname = preview.ApplicantSurname
		previews[i].Image = preview.Image
		previews[i].Title = preview.Title
		previews[i].Location = preview.Location
	}
	sendErr := mh.mailUseCase.SendEmployerMailing(in.Emails, previews)
	if sendErr != nil {
		metrics.MailRequest.WithLabelValues("500", "send message error", "SendEmployerMailing").Inc()
		return &handler.Nothing{}, sendErr
	}
	metrics.MailRequest.WithLabelValues("200", "success", "SendEmployerMailing").Inc()
	return &handler.Nothing{}, nil
}
