package session

import (
	"HeadHunter/auth_microservice/handler"
	"HeadHunter/pkg/errorHandler"
	"context"
	"errors"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type SessionMicroservice struct {
	ctx    context.Context
	client handler.AuthCheckerClient
}

func NewSessionMicroservice(_client handler.AuthCheckerClient) *SessionMicroservice {
	return &SessionMicroservice{
		client: _client,
		ctx:    context.Background(),
	}
}

func (gs *SessionMicroservice) NewSession(email string) (string, error) {
	token, createErr := gs.client.NewSession(gs.ctx, &handler.Email{Value: email})
	if createErr != nil {
		return "", errors.Unwrap(createErr)
	}
	return token.Value, nil
}

func (gs *SessionMicroservice) GetSession(token string) (string, error) {
	email, getErr := gs.client.GetSession(gs.ctx, &handler.Token{Value: token})
	if getErr != nil {
		if errors.Is(getErr, redis.Nil) {
			return "", errorHandler.ErrUnauthorized
		}
		return "", errorHandler.ErrSessionNotFound
	}
	return email.Value, nil
}

func (gs *SessionMicroservice) Delete(token string) error {
	_, deleteErr := gs.client.DeleteSession(gs.ctx, &handler.Token{Value: token})
	if deleteErr != nil {
		return errors.Unwrap(deleteErr)
	}
	return nil
}

func (gs *SessionMicroservice) CreateConfirmationCode(email string) (string, error) {
	code, createErr := gs.client.CreateConfirmationCode(gs.ctx, &handler.Email{Value: email})
	if createErr != nil {
		return "", errors.Unwrap(createErr)
	}
	return code.Value, nil
}

func (gs *SessionMicroservice) GetEmailFromCode(token string) (string, error) {
	email, getErr := gs.client.GetEmailFromCode(gs.ctx, &handler.Token{Value: token})
	if getErr != nil {
		logrus.Println(getErr)
		return "", errorHandler.ErrCodeNotFound
	}
	return email.Value, nil
}

func (gs *SessionMicroservice) GetCodeFromEmail(email string) (string, error) {
	code, getErr := gs.client.GetCodeFromEmail(gs.ctx, &handler.Email{Value: email})
	if getErr != nil {
		return "", errorHandler.ErrCodeNotFound
	}
	return code.Value, nil
}
