package db

import (
	"amarolio-auth/src/customerrors"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheHandler struct {
	rc *redis.Client
	lg Logger
}

func NewCacheHandler(lg Logger) *CacheHandler {
	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("AUTH_REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("AUTH_REDIS_PASSWORD"),
	})
	return &CacheHandler{rc, lg}
}

func (ch *CacheHandler) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := ch.rc.Set(ctx, key, value, expiration).Err(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	ch.lg.Infoln(fmt.Sprintf("SET %s:%s, exp:%d", key, value, expiration))
	return nil
}

func (ch *CacheHandler) Incr(ctx context.Context, key string) (int64, error) {
	val, err := ch.rc.Incr(ctx, key).Result()
	if err != nil {
		return 0, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return val, nil
}

func (ch *CacheHandler) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if err := ch.rc.Expire(ctx, key, expiration).Err(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return nil
}

func (ch *CacheHandler) Del(ctx context.Context, key string) error {
	if err := ch.rc.Del(ctx, key).Err(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return nil
}

func (ch *CacheHandler) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := ch.rc.TTL(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, customerrors.NewError(
				"item not found",
				err,
				customerrors.ItemNotFound,
			)
		}
		return 0, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if ttl < 0 {
		return 0, nil
	}
	return ttl, nil
}

func (ch *CacheHandler) Get(ctx context.Context, key string) (string, error) {
	val, err := ch.rc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", customerrors.NewError(
				"item not found",
				err,
				customerrors.ItemNotFound,
			)
		}
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	ch.lg.Infoln(fmt.Sprintf("GET %s:%s", key, val))
	return val, nil
}
