package repository

//go:generate mockgen -source repository.go -destination=mocks/mock.go

type Repository interface {
	NewSession(email string) (string, error)
	GetSession(token string) (string, error)
	Delete(token string) error
	CreateConfirmationCode(token string) (string, error)
	GetCodeFromEmail(email string) (string, error)
}
