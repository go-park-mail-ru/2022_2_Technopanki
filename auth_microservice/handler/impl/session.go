package impl

import (
	"HeadHunter/auth_microservice/handler"
	"HeadHunter/auth_microservice/usecase/impl"
	"HeadHunter/pkg/errorHandler"
	"context"
)

type SessionHandler struct {
	sessionUseCase impl.SessionUseCase
	handler.UnimplementedAuthCheckerServer
}

func NewSessionHandler(_sus impl.SessionUseCase) *SessionHandler {
	return &SessionHandler{sessionUseCase: _sus}
}

func (sh *SessionHandler) NewSession(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	if in == nil {
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	token, createErr := sh.sessionUseCase.NewSession(in.Value)
	if createErr != nil {
		return &handler.Token{}, createErr
	}
	return &handler.Token{Value: token}, nil
}

func (sh *SessionHandler) GetSession(ctx context.Context, in *handler.Token) (*handler.Email, error) {
	if in == nil {
		return &handler.Email{}, errorHandler.ErrBadRequest
	}
	email, getErr := sh.sessionUseCase.GetSession(in.Value)
	if getErr != nil {
		return &handler.Email{}, getErr
	}
	return &handler.Email{Value: email}, nil
}

func (sh *SessionHandler) DeleteSession(ctx context.Context, in *handler.Token) (*handler.Nothing, error) {
	if in == nil {
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	deleteErr := sh.sessionUseCase.DeleteSession(in.Value)
	if deleteErr != nil {
		return &handler.Nothing{}, deleteErr
	}
	return &handler.Nothing{}, nil
}

func (sh *SessionHandler) CreateConfirmationCode(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	if in == nil {
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	code, createErr := sh.sessionUseCase.CreateConfirmationCode(in.Value)
	if createErr != nil {
		return &handler.Token{}, createErr
	}
	return &handler.Token{Value: code}, nil
}

func (sh *SessionHandler) GetCodeFromEmail(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	if in == nil {
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	code, getErr := sh.sessionUseCase.GetCodeFromEmail(in.Value)
	if getErr != nil {
		return &handler.Token{}, getErr
	}
	return &handler.Token{Value: code}, nil
}
