package impl

import (
	"HeadHunter/auth_microservice/handler"
	"HeadHunter/auth_microservice/usecase"
	"HeadHunter/auth_microservice/usecase/impl"
	"HeadHunter/metrics"
	"HeadHunter/pkg/errorHandler"
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type SessionHandler struct {
	sessionUseCase usecase.Repository
	handler.UnimplementedAuthCheckerServer
}

func NewSessionHandler(_sus *impl.SessionUseCase) *SessionHandler {
	return &SessionHandler{sessionUseCase: _sus}
}

func (sh *SessionHandler) NewSession(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	timer := prometheus.NewTimer(metrics.SessionRequestDuration.WithLabelValues("NewSession"))
	defer timer.ObserveDuration()

	if in == nil {
		metrics.SessionRequest.WithLabelValues("400", "bad request NewSession")
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	token, createErr := sh.sessionUseCase.NewSession(in.Value)
	if createErr != nil {
		metrics.SessionRequest.WithLabelValues("500", "cannot create session")
		return &handler.Token{}, createErr
	}
	metrics.SessionRequest.WithLabelValues("200", "success NewSession")
	return &handler.Token{Value: token}, nil
}

func (sh *SessionHandler) GetSession(ctx context.Context, in *handler.Token) (*handler.Email, error) {
	timer := prometheus.NewTimer(metrics.SessionRequestDuration.WithLabelValues("GetSession"))
	defer timer.ObserveDuration()

	if in == nil {
		metrics.SessionRequest.WithLabelValues("400", "bad request GetSession")
		return &handler.Email{}, errorHandler.ErrBadRequest
	}
	email, getErr := sh.sessionUseCase.GetSession(in.Value)
	if getErr != nil {
		metrics.SessionRequest.WithLabelValues("404", "session not found")
		return &handler.Email{}, getErr
	}
	metrics.SessionRequest.WithLabelValues("200", "success GetSession")
	return &handler.Email{Value: email}, nil
}

func (sh *SessionHandler) DeleteSession(ctx context.Context, in *handler.Token) (*handler.Nothing, error) {
	timer := prometheus.NewTimer(metrics.SessionRequestDuration.WithLabelValues("DeleteSession"))
	defer timer.ObserveDuration()

	if in == nil {
		metrics.SessionRequest.WithLabelValues("400", "bad request DeleteSession")
		return &handler.Nothing{}, errorHandler.ErrBadRequest
	}
	deleteErr := sh.sessionUseCase.DeleteSession(in.Value)
	if deleteErr != nil {
		metrics.SessionRequest.WithLabelValues("500", "cannot delete session")
		return &handler.Nothing{}, deleteErr
	}
	metrics.SessionRequest.WithLabelValues("200", "success DeleteSession")
	return &handler.Nothing{}, nil
}

func (sh *SessionHandler) CreateConfirmationCode(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	timer := prometheus.NewTimer(metrics.SessionRequestDuration.WithLabelValues("CreateConfirmationCode"))
	defer timer.ObserveDuration()

	if in == nil {
		metrics.SessionRequest.WithLabelValues("400", "bad request CreateConfirmationCode")
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	code, createErr := sh.sessionUseCase.CreateConfirmationCode(in.Value)
	if createErr != nil {
		metrics.SessionRequest.WithLabelValues("500", "cannot create confirmation code")
		return &handler.Token{}, createErr
	}
	metrics.SessionRequest.WithLabelValues("200", "success CreateConfirmationCode")
	return &handler.Token{Value: code}, nil
}

func (sh *SessionHandler) GetCodeFromEmail(ctx context.Context, in *handler.Email) (*handler.Token, error) {
	timer := prometheus.NewTimer(metrics.SessionRequestDuration.WithLabelValues("GetCodeFromEmail"))
	defer timer.ObserveDuration()

	if in == nil {
		metrics.SessionRequest.WithLabelValues("400", "bad request GetCodeFromEmail")
		return &handler.Token{}, errorHandler.ErrBadRequest
	}
	code, getErr := sh.sessionUseCase.GetCodeFromEmail(in.Value)
	if getErr != nil {
		metrics.SessionRequest.WithLabelValues("404", "code not found")
		return &handler.Token{}, getErr
	}
	metrics.SessionRequest.WithLabelValues("200", "success GetCodeFromEmail")
	return &handler.Token{Value: code}, nil
}
