package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("AUTH_REDIS_HOST") + ":" + os.Getenv("AUTH_REDIS_PORT"),
		Password: os.Getenv("AUTH_REDIS_PASS"),
	})
}
