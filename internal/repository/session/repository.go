package session

type Token string
type Session struct {
	Email     string
	ExpiresAt int64
}

type Repository interface {
	NewSession(email string) (string, error)
	GetSession(token Token) (string, error)
	DeleteSession(token Token) error
	Expiring() int64
}
