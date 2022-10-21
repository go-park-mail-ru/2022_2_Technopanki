package session

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"github.com/google/uuid"
	"sync"
	"time"
)

func (s *Session) IsExpired() bool {
	return time.Now().Unix() > s.ExpiresAt
}

type Store struct {
	Values           map[Token]Session
	DefaultExpiresAt time.Duration
	mutex            sync.RWMutex
}

func (s *Store) Expiring() time.Duration {
	return s.DefaultExpiresAt
}

func (s *Store) NewSession(email string) string {
	token := uuid.NewString()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Values[Token(token)] = Session{
		Email:     email,
		ExpiresAt: time.Now().Unix() + int64(s.DefaultExpiresAt),
	}

	return token
}

func (s *Store) GetSession(token Token) (Session, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if val, ok := s.Values[token]; ok {
		return val, nil
	}

	return Session{}, errorHandler.ErrSessionNotFound
}

func (s *Store) DeleteSession(token Token) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.Values[token]; ok {
		delete(s.Values, token)
		return nil
	}

	return errorHandler.ErrCannotDeleteSession
}

func NewSessionsStore(cfg configs.Config) *Store {
	return &Store{
		Values:           make(map[Token]Session),
		DefaultExpiresAt: time.Duration(cfg.DefaultExpiringSession) * time.Hour / time.Second,
	}
}
