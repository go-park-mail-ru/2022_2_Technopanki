package session

type Session struct {
	Email     string
	ExpiresAt int
}

//go:generate mockgen -source repository.go -destination=mocks/mock.go

type Repository interface {
	NewSession(email string) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}
