package session

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RedisStore struct {
	DefaultExpiresAt int
	client           *redis.Client
}

func NewRedisStore(cfg configs.Config, _redis *redis.Client) *RedisStore {
	return &RedisStore{
		client:           _redis,
		DefaultExpiresAt: cfg.DefaultExpiringSession,
	}
}

func (rs *RedisStore) NewSession(email string) (string, error) {
	token := uuid.NewString()
	err := rs.client.Do("SETEX", token, rs.DefaultExpiresAt, email).Err()
	if err != nil {
		return "", fmt.Errorf("creating session error: %w", err)
	}
	return token, nil
}

func (rs *RedisStore) GetSession(token string) (string, error) {
	result, getErr := rs.client.Get(token).Result()
	if getErr != nil {
		return "", fmt.Errorf("getting session error: %w", getErr)
	}
	return result, nil
}

func (rs *RedisStore) DeleteSession(token string) error {
	err := rs.client.Del(token).Err()
	if err != nil {
		return fmt.Errorf("deleting session error: %w", err)
	}
	return nil
}
