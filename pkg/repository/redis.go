package repository

import (
	"HeadHunter/configs"
	"fmt"
	"github.com/go-redis/redis"
)

func RedisConnect(cfg configs.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})

	redisErr := client.Ping().Err()
	if redisErr != nil {
		return &redis.Client{}, redisErr
	}
	return client, nil
}
