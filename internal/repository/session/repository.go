package session

type Session struct {
	Email     string
	ExpiresAt int
}

type Repository interface {
	NewSession(email string) (string, error)
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}
