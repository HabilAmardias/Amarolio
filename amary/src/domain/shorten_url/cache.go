package shortenurl

import (
	"amary/src/customerror"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TaggedShortenUrl struct {
	ID               int64      `json:"id"`
	UserID           *string    `json:"user_id"`
	EncryptedLongUrl string     `json:"encrypted_long_url"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	ExpiredAt        *time.Time `json:"expired_at"`
}

type ShortenURLCache struct {
	rc *redis.Client
}

func NewShortenURLCache(rc *redis.Client) *ShortenURLCache {
	return &ShortenURLCache{rc}
}

func (suc *ShortenURLCache) AddToUserSet(ctx context.Context, userID string, cat time.Time, encodedID string) error {
	key := fmt.Sprintf("user:%s:list:shorten_url", userID)
	item := redis.Z{
		Score:  float64(cat.Unix()),
		Member: encodedID,
	}
	if err := suc.rc.ZAdd(ctx, key, item).Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return nil
}

func (suc *ShortenURLCache) Set(ctx context.Context, encodedID string, ttl time.Duration, url ShortenURL) error {
	taggedUrl := TaggedShortenUrl(url)
	b, err := json.Marshal(taggedUrl)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("shorten_url:%s", encodedID)
	_, err = suc.rc.Set(ctx, key, string(b), ttl).Result()
	return err
}

func (suc *ShortenURLCache) Get(ctx context.Context, encodedID string, url *ShortenURL) error {
	key := fmt.Sprintf("shorten_url:%s", encodedID)
	taggedUrl := new(TaggedShortenUrl)

	val, err := suc.rc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerror.NewError(
				"url not found",
				err,
				customerror.ItemNotFound,
			)
		}
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	if err := json.Unmarshal([]byte(val), taggedUrl); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}

	*url = ShortenURL(*taggedUrl)
	return nil
}
