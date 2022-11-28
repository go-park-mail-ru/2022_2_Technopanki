package session

import (
	"HeadHunter/auth_microservice/handler"
	"context"
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
		return "", createErr
	}
	return token.Value, nil
}

func (gs *SessionMicroservice) GetSession(token string) (string, error) {
	email, getErr := gs.client.GetSession(gs.ctx, &handler.Token{Value: token})
	if getErr != nil {
		return "", getErr
	}
	return email.Value, nil
}

func (gs *SessionMicroservice) DeleteSession(token string) error {
	_, deleteErr := gs.client.DeleteSession(gs.ctx, &handler.Token{Value: token})
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (gs *SessionMicroservice) CreateConfirmationCode(email string) (string, error) {
	code, createErr := gs.client.CreateConfirmationCode(gs.ctx, &handler.Email{Value: email})
	if createErr != nil {
		return "", createErr
	}
	return code.Value, nil
}

func (gs *SessionMicroservice) GetEmailFromCode(token string) (string, error) {
	email, getErr := gs.client.GetEmailFromCode(gs.ctx, &handler.Token{Value: token})
	if getErr != nil {
		return "", getErr
	}
	return email.Value, nil
}
