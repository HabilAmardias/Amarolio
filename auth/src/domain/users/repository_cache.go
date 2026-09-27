package users

import (
	"amarolio-auth/src/customerrors"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CacheHandlerItf interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type UserCacheImpl struct {
	ch CacheHandlerItf
}

func NewUserCache(rc CacheHandlerItf) *UserCacheImpl {
	return &UserCacheImpl{rc}
}

func (uc *UserCacheImpl) SetCacheByEmail(ctx context.Context, age time.Duration, user *User) error {
	key := fmt.Sprintf("users:email:%s", user.Email)
	val, err := json.Marshal(*user)
	if err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if err := uc.ch.Set(ctx, key, string(val), age); err != nil {
		return err
	}
	return nil
}

func (uc *UserCacheImpl) FindCacheByEmail(ctx context.Context, userEmail string, user *User) error {
	key := fmt.Sprintf("users:email:%s", userEmail)
	val, err := uc.ch.Get(ctx, key)
	if err != nil {
		return err
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

func (uc *UserCacheImpl) SetCacheByID(ctx context.Context, age time.Duration, user *User) error {
	key := fmt.Sprintf("users:id:%s", user.ID)
	val, err := json.Marshal(*user)
	if err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return uc.ch.Set(ctx, key, string(val), age)
}

func (uc *UserCacheImpl) FindCacheByID(ctx context.Context, userID string, user *User) error {
	key := fmt.Sprintf("users:id:%s", userID)
	val, err := uc.ch.Get(ctx, key)
	if err != nil {
		return err
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
