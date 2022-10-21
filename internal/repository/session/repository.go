package session

import "time"

type Token string
type Session struct {
	Email     string
	ExpiresAt int64
}

type Repository interface {
	NewSession(email string) string
	GetSession(token Token) (Session, error)
	DeleteSession(token Token) error
	Expiring() time.Duration
}
