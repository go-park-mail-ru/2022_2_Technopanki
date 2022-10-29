package session

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RedisStore struct {
	DefaultExpiresAt int64
	client           *redis.Client
}

func NewRedisStore(cfg configs.Config, _redis *redis.Client) *RedisStore {
	return &RedisStore{
		client:           _redis,
		DefaultExpiresAt: cfg.DefaultExpiringSession,
	}
}

func (rs *RedisStore) Expiring() int64 {
	return rs.DefaultExpiresAt
}

func (rs *RedisStore) NewSession(email string) (string, error) {
	token := uuid.NewString()
	err := rs.client.Do("SETEX", token, rs.DefaultExpiresAt, email).Err()
	if err != nil {
		return "", errorHandler.ErrCannotCreateSession
	}
	return token, nil
}

func (rs *RedisStore) GetSession(token Token) (string, error) {
	result, getErr := rs.client.Get(string(token)).Result()
	if getErr != nil {
		return "", errorHandler.ErrSessionNotFound
	}
	return result, nil
}

func (rs *RedisStore) DeleteSession(token Token) error {
	err := rs.client.Del(string(token)).Err()
	if err != nil {
		return errorHandler.ErrCannotDeleteSession
	}
	return nil
}
