package sessions

import (
	"time"
)

type Session struct {
	Email     string
	ExpiresAt int64
}

func (s *Session) IsExpired() bool {
	return time.Now().Unix() > s.ExpiresAt
}
