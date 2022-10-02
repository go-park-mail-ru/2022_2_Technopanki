package sessions

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Token string

type Store struct {
	Values           map[Token]Session
	DefaultExpiresAt int64
	mutex            sync.RWMutex
}

func NewStore() Store {
	return Store{
		Values: make(map[Token]Session),
	}
}

func (s *Store) NewSession(email string) string {
	token := uuid.NewString()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Values[Token(token)] = Session{
		Email:     email,
		ExpiresAt: time.Now().Unix() + s.DefaultExpiresAt,
	}

	return token
}

func (s *Store) GetSession(token Token) (Session, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if val, ok := s.Values[token]; ok {
		return val, nil
	}

	return Session{}, errors.New("no session with this token")
}

var SessionsStore = NewStore()
