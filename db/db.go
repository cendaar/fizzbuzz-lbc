package db

import (
	"github.com/go-redis/redis"
	"os"
)

type RedisInstance struct {
	Client *redis.Client
}

func Initialize() (*RedisInstance, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisInstance{Client: client}, nil
}