package users

import (
	"amarolio-auth/src/customerrors"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CacheHandlerItf interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	TTL(ctx context.Context, key string) (time.Duration, error)
}

func loginAttemptsKey(email string) string {
	return fmt.Sprintf("auth:attempts:%s", strings.ToLower(strings.TrimSpace(email)))
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

func (uc *UserCacheImpl) GetLoginAttempts(ctx context.Context, email string) (int, error) {
	val, err := uc.ch.Get(ctx, loginAttemptsKey(email))
	if err != nil {
		var ce *customerrors.CustomError
		if errors.As(err, &ce) && ce.ErrCode == customerrors.ItemNotFound {
			return 0, nil
		}
		return 0, err
	}
	attempts, convErr := strconv.Atoi(val)
	if convErr != nil {
		return 0, customerrors.NewError(
			"something went wrong",
			convErr,
			customerrors.CommonErr,
		)
	}
	return attempts, nil
}

// IncreaseLoginAttempts increments the failed-attempt counter. The first
// failure starts the window; reaching the max re-arms it so the lockout lasts
// a full `lock` duration from the final failure.
func (uc *UserCacheImpl) IncreaseLoginAttempts(ctx context.Context, email string, max int, lock time.Duration) (int, error) {
	key := loginAttemptsKey(email)
	count, err := uc.ch.Incr(ctx, key)
	if err != nil {
		return 0, err
	}
	if count == 1 || int(count) >= max {
		if err := uc.ch.Expire(ctx, key, lock); err != nil {
			return int(count), err
		}
	}
	return int(count), nil
}

func (uc *UserCacheImpl) ResetLoginAttempts(ctx context.Context, email string) error {
	return uc.ch.Del(ctx, loginAttemptsKey(email))
}

func (uc *UserCacheImpl) LoginLockRemaining(ctx context.Context, email string) (time.Duration, error) {
	return uc.ch.TTL(ctx, loginAttemptsKey(email))
}
