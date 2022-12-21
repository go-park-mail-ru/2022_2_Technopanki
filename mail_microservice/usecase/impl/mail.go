package impl

import (
	"HeadHunter/common/session"
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/usecase/sender"
	"HeadHunter/pkg/errorHandler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type MailService struct {
	sessionRepo session.Repository
	sender      sender.Sender
}

func NewMailService(_sessionRepo session.Repository, _sender sender.Sender) *MailService {
	return &MailService{sessionRepo: _sessionRepo, sender: _sender}
}

func (ms *MailService) SendConfirmCode(email string) error {
	_, getErr := ms.sessionRepo.GetCodeFromEmail(email)
	if getErr == nil {
		return status.Error(codes.AlreadyExists, errorHandler.ErrCodeAlreadyExists.Error())
	}

	code, createErr := ms.sessionRepo.CreateConfirmationCode(email)
	if createErr != nil {
		return createErr
	}
	sendErr := ms.sender.SendConfirmCode(email, code)
	if sendErr != nil {
		return sendErr
	}
	return nil
}

func (ms *MailService) SendApplicantMailing(emails []string, vacancies []*models.VacancyPreview) error {
	vacanciesObjects := make([]models.VacancyPreview, len(vacancies))
	for i, vacancy := range vacancies {
		vacanciesObjects[i] = *vacancy
	}

	for _, email := range emails {
		err := ms.sender.SendApplicantMailing(email, vacanciesObjects)
		if err != nil {
			log.Printf("error with send to %s: %s", email, err)
		}
	}
	return nil
}

func (ms *MailService) SendEmployerMailing(emails []string, previews []*models.ResumePreview) error {
	previewObject := make([]models.ResumePreview, len(previews))
	log.Println("len(previews):", len(previews))
	for i, applicant := range previews {
		previewObject[i] = *applicant
		log.Println(previewObject[i])
	}

	for _, email := range emails {
		err := ms.sender.SendEmployerMailing(email, previewObject)
		if err != nil {
			log.Printf("error with send to %s: %s", email, err)
		}
	}
	return nil
}
