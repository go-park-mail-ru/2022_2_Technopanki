package mail

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases/sender"
	"fmt"
)

type MailService struct {
	userRepo    repository.UserRepository
	sessionRepo session.Repository
	sender      sender.Sender
}

func NewMailService(repo repository.UserRepository, _sessionRepo session.Repository, _sender sender.Sender) *MailService {
	return &MailService{userRepo: repo, sessionRepo: _sessionRepo, sender: _sender}
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

func (ms *MailService) SendApplicantMailing(users []*models.UserAccount, vacancies []*models.Vacancy) error {
	var returnedErr error
	for _, user := range users {
		err := ms.sender.SendApplicantMailing(user.Email, vacancies)
		if err != nil {
			returnedErr = fmt.Errorf("%s %w", err.Error(), returnedErr)
		}
	}
	return returnedErr
}

func (ms *MailService) SendEmployerMailing(employers, applicants []*models.UserAccount) error {
	var returnedErr error
	for _, user := range employers {
		err := ms.sender.SendEmployerMailing(user.Email, applicants)
		if err != nil {
			returnedErr = fmt.Errorf("%s %w", err.Error(), returnedErr)
		}
	}
	return returnedErr
}
