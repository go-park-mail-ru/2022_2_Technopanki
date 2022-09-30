package handlers

import (
	"github.com/google/uuid"
	"time"
)

type sessionElement struct {
	email     string
	expiresAt time.Time
}

func (s *sessionElement) isExpired() bool {
	return time.Now().Unix() < s.expiresAt.Unix()
}

type sessions struct {
	sessionsMap map[string]sessionElement
}

func (s *sessions) NewSession(email string) string {
	token := uuid.NewString()
	s.sessionsMap[token] = sessionElement{
		email:     email,
		expiresAt: time.Now().Add(time.Minute * 5),
	}

	return token
}

var SessionsMap = sessions{
	sessionsMap: make(map[string]sessionElement, 0),
}
