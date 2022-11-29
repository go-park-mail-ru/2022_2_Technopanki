package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/mail_microservice/repository/session"
	"HeadHunter/mail_microservice/usecase/sender"
)

type MailService struct {
	sessionRepo session.Repository
	sender      sender.Sender
}

func NewMailService(_sessionRepo session.Repository, _sender sender.Sender) *MailService {
	return &MailService{sessionRepo: _sessionRepo, sender: _sender}
}

func (ms *MailService) SendConfirmCode(email string) error {
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

func (ms *MailService) SendApplicantMailing(emails []string, vacancies []*models.Vacancy) error {
	for _, email := range emails {
		err := ms.sender.SendApplicantMailing(email, vacancies)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ms *MailService) SendEmployerMailing(emails []string, applicants []*models.UserAccount) error {
	for _, email := range emails {
		err := ms.sender.SendEmployerMailing(email, applicants)
		if err != nil {
			return err
		}
	}
	return nil
}