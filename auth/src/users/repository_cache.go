package users

import (
	"amarolio-auth/src/customerrors"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserCacheImpl struct {
	rc  *redis.Client
	age time.Duration
}

func NewUserCache(rc *redis.Client) *UserCacheImpl {
	// no expiration
	age := 0 * time.Second
	return &UserCacheImpl{rc, age}
}

func (uc *UserCacheImpl) SetCacheByEmail(ctx context.Context, user *User) error {
	key := fmt.Sprintf("users:email:%s", user.Email)
	val, err := json.Marshal(*user)
	if err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if err := uc.rc.Set(ctx, key, string(val), uc.age).Err(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (uc *UserCacheImpl) FindCacheByEmail(ctx context.Context, userEmail string, user *User) error {
	key := fmt.Sprintf("users:email:%s", userEmail)
	val, err := uc.rc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerrors.NewError(
				"no user found",
				err,
				customerrors.ItemNotFound,
			)
		}
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	if err := json.Unmarshal([]byte(val), user); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return nil
}

func (uc *UserCacheImpl) SetCacheByID(ctx context.Context, user *User) error {
	key := fmt.Sprintf("users:id:%s", user.ID)
	val, err := json.Marshal(*user)
	if err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if err := uc.rc.Set(ctx, key, string(val), uc.age).Err(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (uc *UserCacheImpl) FindCacheByID(ctx context.Context, userID string, user *User) error {
	key := fmt.Sprintf("users:id:%s", userID)
	val, err := uc.rc.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerrors.NewError(
				"no user found",
				err,
				customerrors.ItemNotFound,
			)
		}
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	if err := json.Unmarshal([]byte(val), user); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return nil
}
