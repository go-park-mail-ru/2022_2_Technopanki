package impl

import "HeadHunter/auth_microservice/repository/impl"

type SessionUseCase struct {
	redisStore impl.RedisStore
}

func NewSessionUseCase(_rs impl.RedisStore) *SessionUseCase {
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

func (sus *SessionUseCase) GetEmailFromCode(token string) (string, error) {
	return sus.redisStore.GetEmailFromCode(token)
}
