package repository

import (
	"HeadHunter/configs"
	"fmt"
	"github.com/go-redis/redis"
)

func RedisConnect(cfg configs.RedisConfig) (*redis.Client, error) {
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Password,
		DB:       0,
	})

	redisErr := client.Ping().Err()
	if redisErr != nil {
		return &redis.Client{}, redisErr
	}
	return client, nil
}
