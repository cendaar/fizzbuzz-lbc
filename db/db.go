package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

type RedisInstance struct {
	Client *redis.Client
}

func Initialize() (*RedisInstance, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisInstance{Client: client}, nil
}