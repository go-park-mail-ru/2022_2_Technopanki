package session

import (
	"HeadHunter/configs"
	"HeadHunter/internal/errorHandler"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

type RedisStore struct {
	DefaultExpiresAt time.Duration
	client           *redis.Client
}

func NewRedisStore(cfg configs.Config, _redis *redis.Client) *RedisStore {
	return &RedisStore{
		client:           _redis,
		DefaultExpiresAt: time.Duration(cfg.DefaultExpiringSession) * time.Hour / time.Second,
	}
}

func (rs *RedisStore) Expiring() time.Duration {
	return rs.DefaultExpiresAt
}

func (rs *RedisStore) NewSession(email string) (string, error) {
	ctx := context.Background()
	token := uuid.NewString()
	err := rs.client.Set(ctx, token, email, rs.DefaultExpiresAt).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (rs *RedisStore) GetSession(token Token) (string, error) {
	ctx := context.Background()
	result, getErr := rs.client.Get(ctx, string(token)).Result()
	if getErr != nil {
		return "", errorHandler.ErrSessionNotFound
	}
	return result, nil
}

func (rs *RedisStore) DeleteSession(token Token) error {
	ctx := context.Background()
	err := rs.client.Del(ctx, string(token)).Err()
	if err != nil {
		return errorHandler.ErrCannotDeleteSession
	}
	return nil
}
