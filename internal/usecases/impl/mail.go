package impl

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

func (ms *MailService) ConfirmationAccount(email string) error {
	token, tokenErr := ms.sessionRepo.CreateConfirmationToken(email)
	if tokenErr != nil {
		return tokenErr
	}
	sendErr := ms.sender.SendConfirmToken(email, token)
	if sendErr != nil {
		return sendErr
	}
	return nil
}

func (ms *MailService) UpdatePassword() {

}

func (ms *MailService) TwoFactorSignIn() {

}
