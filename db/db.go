package db

import (
	"github.com/go-redis/redis"
	"log"
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

	log.Println(os.Getenv("REDIS_URL"))

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisInstance{Client: client}, nil
}