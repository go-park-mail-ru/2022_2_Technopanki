package sessions

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Token string

type Store struct {
	Values           map[Token]Session
	DefaultExpiresAt time.Duration
	mutex            sync.RWMutex
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

var SessionsStore Store

//var SessionsStore = Store{
//	Values:           make(map[Token]Session),
//	DefaultExpiresAt: 12 * time.Hour / time.Second,
//}

func NewSessionsStore(cfg configs.Config) *Store {
	return &Store{
		Values:           make(map[Token]Session),
		DefaultExpiresAt: time.Duration(cfg.DefaultExpiringSession) * time.Hour / time.Second,
	}
}
