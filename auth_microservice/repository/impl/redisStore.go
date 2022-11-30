package impl

import (
	"HeadHunter/auth_microservice/configs"
	"HeadHunter/auth_microservice/usecase/codeGenerator"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RedisStore struct {
	DefaultExpiresAt int
	ConfirmationTime int
	client           *redis.Client
}

func NewRedisStore(cfg *configs.SessionConfig, _redis *redis.Client) *RedisStore {
	return &RedisStore{
		client:           _redis,
		DefaultExpiresAt: cfg.DefaultExpiringSession,
		ConfirmationTime: cfg.ConfirmationTime,
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

func (rs *RedisStore) Delete(token string) error {
	err := rs.client.Del(token).Err()
	if err != nil {
		return fmt.Errorf("deleting error: %w", err)
	}
	return nil
}

func (rs *RedisStore) CreateConfirmationCode(email string) (string, error) {
	var code string
	var generateErr error

	for {
		code, generateErr = codeGenerator.GenerateCode()
		if generateErr != nil {
			return "", generateErr
		}
		_, getErr := rs.client.Get(code).Result()
		if getErr != nil {
			break
		}
	}

	err := rs.client.Do("SETEX", email, rs.ConfirmationTime, code).Err()
	if err != nil {
		return "", fmt.Errorf("creating confirmation code error: %w", err)
	}
	return code, nil
}

func (rs *RedisStore) GetEmailFromCode(token string) (string, error) {
	result, getErr := rs.client.Get(token).Result()
	if getErr != nil {
		return "", fmt.Errorf("getting code error: %w", getErr)
	}

	return result, nil
}

func (rs *RedisStore) GetCodeFromEmail(email string) (string, error) {
	result, getErr := rs.client.Get(email).Result()
	if getErr != nil {
		return "", fmt.Errorf("getting code error: %w", getErr)
	}

	return result, nil
}
