package url

import (
	"amary/src/customerror"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TaggedURL struct {
	ID               int64      `json:"id"`
	UserID           *string    `json:"user_id"`
	EncryptedLongUrl string     `json:"encrypted_long_url"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	ExpiredAt        *time.Time `json:"expired_at"`
}

type URLCache struct {
	rc *redis.Client
}

func NewShortenURLCache(rc *redis.Client) *URLCache {
	return &URLCache{rc}
}

func (suc *URLCache) GetUserLinks(ctx context.Context, ids *[]string, links *[]URL, missingIDs *[]string) error {
	keys := []string{}
	for i, id := range *ids {
		keys[i] = fmt.Sprintf("shorten_url:%s", id)
	}
	res, err := suc.rc.MGet(ctx, keys...).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerror.NewError(
				"item not found",
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
	if len(res) == 0 {
		return customerror.NewError(
			"item not found",
			err,
			customerror.ItemNotFound,
		)
	}
	for i, val := range res {
		if val == nil {
			*missingIDs = append(*missingIDs, (*ids)[i])
			continue
		}
		var l TaggedURL
		if err := json.Unmarshal([]byte(val.(string)), &l); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.CommonErr,
			)
		}
		*links = append(*links, URL(l))
	}

	return nil
}

func (suc *URLCache) GetIDFromUserSet(ctx context.Context, userID string, page int64, limit int64, ids *[]string) error {
	key := fmt.Sprintf("user:%s:list:shorten_url", userID)
	start := (page - 1) * limit
	stop := start + limit - 1
	res, err := suc.rc.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   key,
		Start: start,
		Stop:  stop,
		Rev:   true,
	}).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerror.NewError(
				"item not found",
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
	if len(res) == 0 {
		return customerror.NewError(
			"item not found",
			errors.New("item not found"),
			customerror.ItemNotFound,
		)
	}
	*ids = res
	return nil
}

func (suc *URLCache) AddToUserSet(ctx context.Context, userID string, cat time.Time, encodedID string) error {
	key := fmt.Sprintf("user:%s:list:shorten_url", userID)
	item := redis.Z{
		Score:  float64(cat.Unix()),
		Member: encodedID,
	}
	pipe := suc.rc.Pipeline()
	pipe.ZAdd(ctx, key, item)
	pipe.Expire(ctx, key, 24*time.Hour)
	if _, err := pipe.Exec(ctx); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return nil
}

func (suc *URLCache) Set(ctx context.Context, encodedID string, ttl time.Duration, url URL) error {
	taggedUrl := TaggedURL(url)
	b, err := json.Marshal(taggedUrl)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("shorten_url:%s", encodedID)
	_, err = suc.rc.Set(ctx, key, string(b), ttl).Result()
	return err
}

func (suc *URLCache) Get(ctx context.Context, encodedID string, url *URL) error {
	key := fmt.Sprintf("shorten_url:%s", encodedID)
	taggedUrl := new(TaggedURL)

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

	*url = URL(*taggedUrl)
	return nil
}
