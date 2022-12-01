package impl

import (
	"HeadHunter/auth_microservice/repository"
	"HeadHunter/common/session"
)

type SessionUseCase struct {
	redisStore repository.Repository
}

func NewSessionUseCase(_rs session.Repository) *SessionUseCase {
	return &SessionUseCase{redisStore: _rs}
}

func (sus *SessionUseCase) NewSession(email string) (string, error) {
	return sus.redisStore.NewSession(email)
}

func (sus *SessionUseCase) GetSession(token string) (string, error) {
	return sus.redisStore.GetSession(token)
}

func (sus *SessionUseCase) DeleteSession(token string) error {
	return sus.redisStore.Delete(token)
}

func (sus *SessionUseCase) CreateConfirmationCode(email string) (string, error) {
	return sus.redisStore.CreateConfirmationCode(email)
}

func (sus *SessionUseCase) GetCodeFromEmail(email string) (string, error) {
	return sus.redisStore.GetCodeFromEmail(email)
}
