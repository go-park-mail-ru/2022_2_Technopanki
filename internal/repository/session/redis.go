package session

import (
	"HeadHunter/configs"
	"HeadHunter/pkg/errorHandler"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"strings"
)

type RedisStore struct {
	DefaultExpiresAt int
	ConfirmationTime int
	client           *redis.Client
}

func NewRedisStore(cfg *configs.Config, _redis *redis.Client) *RedisStore {
	return &RedisStore{
		client:           _redis,
		DefaultExpiresAt: cfg.DefaultExpiringSession,
		ConfirmationTime: cfg.Security.ConfirmationTime,
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

func (rs *RedisStore) CreateConfirmationToken(email string) (string, error) {
	token := uuid.NewString()
	err := rs.client.Do("SETEX", token, rs.ConfirmationTime, "confirm:"+email).Err()
	if err != nil {
		return "", fmt.Errorf("creating confirmation token error: %w", err)
	}
	return token, nil
}

func (rs *RedisStore) GetEmailFromConfirmationToken(token string) (string, error) {
	result, getErr := rs.client.Get(token).Result()
	if getErr != nil {
		return "", fmt.Errorf("getting session error: %w", getErr)
	}
	splitResult := strings.Split(result, ":")
	if len(splitResult) != 2 {
		return "", errorHandler.ErrBadRequest
	}
	return splitResult[1], nil
}
