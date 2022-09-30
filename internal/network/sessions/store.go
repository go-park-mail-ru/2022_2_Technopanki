package sessions

import (
	"github.com/google/uuid"
	"time"
)

type Token string

type Store struct {
	Values           map[Token]Session
	DefaultExpiresAt int64
}

func NewStore() Store {
	return Store{
		Values: make(map[Token]Session),
	}
}

func (s *Store) NewSession(email string) string {
	token := uuid.NewString()
	s.Values[Token(token)] = Session{
		Email:     email,
		ExpiresAt: time.Now().Unix() + s.DefaultExpiresAt,
	}

	return token
}

func (s *Store) GetSession(token Token) Session {
	return s.Values[token]
}

func (s *Store) GetToken(email string) Token {
	for k, v := range s.Values {
		if v.Email == email {
			return k
		}
	}

	return ""
}

func (s *Store) UpdateSession(token Token) string {
	user := s.GetSession(token)
	delete(s.Values, token)
	return s.NewSession(user.Email)
}

var SessionsStore = NewStore()

// TODO
//func clearStore() {
//	for key, value := range store.Values {
//		if value.isExpired() {
//			delete(store.Values, key)
//		}
//	}
//}
