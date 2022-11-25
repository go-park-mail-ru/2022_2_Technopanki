package mail

import (
	"HeadHunter/internal/repository"
	"HeadHunter/internal/repository/session"
	"HeadHunter/internal/usecases/sender"
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

func (ms *MailService) UpdatePassword() {

}
