package usecase

//go:generate mockgen -source usecase.go -destination=mocks/mock.go

type Repository interface {
	NewSession(email string) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
	CreateConfirmationCode(email string) (string, error)
	GetCodeFromEmail(email string) (string, error)
}
