package shortenurl

import (
	"amary/src/customerror"
	"amary/src/services"
	"context"
	"errors"
	"time"
)

type URLEncryptorItf interface {
	DecryptURL(cipherURL string) (string, error)
	EncryptURL(plainURL string) string
}

type IDEncoderItf interface {
	Encode(id int64) string
	Decode(encodedID string) (int64, error)
}

type ShortenURLCacheItf interface {
	Set(ctx context.Context, encodedID string, ttl time.Duration, url ShortenURL) error
	Get(ctx context.Context, encodedID string, url *ShortenURL) error
	AddToUserSet(ctx context.Context, userID string, cat time.Time, encodedID string) error
}

type VisitRecordRepoItf interface {
	InsertNewRecord(
		ctx context.Context,
		userID string,
		id int64,
		device string,
	) error
}

type ShortenURLRepoItf interface {
	InsertNewURL(
		ctx context.Context,
		userID *string,
		encryptedLongURL string,
		expiredAt *time.Time,
		shortenURL *ShortenURL,
	) error
	FindByID(ctx context.Context, id int64, url *ShortenURL) error
}

type ShortenURLServiceImpl struct {
	ue  URLEncryptorItf
	ide IDEncoderItf
	suc ShortenURLCacheItf
	sur ShortenURLRepoItf
	vrr VisitRecordRepoItf
}

func NewShortenURLServ(ue URLEncryptorItf, ide IDEncoderItf, suc ShortenURLCacheItf, sur ShortenURLRepoItf, vrr VisitRecordRepoItf) *ShortenURLServiceImpl {
	return &ShortenURLServiceImpl{ue, ide, suc, sur, vrr}
}

func (sus *ShortenURLServiceImpl) NewShortURL(ctx context.Context, userID *string, longURL string, duration *int) (string, error) {
	// default using 24 hour duration
	now := time.Now()
	eatv := now.Add(24 * time.Hour)
	eat := &eatv

	// user logged-in and no expiration time
	if userID != nil && duration == nil {
		eat = nil
	}

	// user logged-in and set expiration time
	if userID != nil && duration != nil {
		expiration := now.Add(time.Duration(*duration) * 24 * time.Hour)
		eat = &expiration
	}

	url := new(ShortenURL)

	encryptedURL := sus.ue.EncryptURL(longURL)
	if err := sus.sur.InsertNewURL(ctx, userID, encryptedURL, eat, url); err != nil {
		return "", err
	}

	encodedID := sus.ide.Encode(url.ID)

	go func(uid *string, eid string, u ShortenURL) {
		ttl := 24 * time.Hour
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// add to cache
		fun := func() error {
			return sus.suc.Set(ctx, eid, ttl, u)
		}

		services.WithErrorRetry(ctx, fun, 100*time.Millisecond)

		// if user is logged-in then add the id to zset
		if uid != nil {
			fun := func() error {
				return sus.suc.AddToUserSet(ctx, *uid, u.CreatedAt, eid)
			}

			services.WithErrorRetry(ctx, fun, 100*time.Millisecond)
		}
	}(userID, encodedID, *url)

	return encodedID, nil
}

func (sus *ShortenURLServiceImpl) FindLongURL(ctx context.Context, encodedID string, device string) (string, error) {
	url := new(ShortenURL)
	now := time.Now()

	decodedID, err := sus.ide.Decode(encodedID)
	if err != nil {
		return "", err
	}

	if err := sus.suc.Get(ctx, encodedID, url); err != nil {
		var ce *customerror.CustomError
		if !errors.As(err, &ce) {
			return "", customerror.NewError(
				"something went wrong",
				errors.New("parse error fail"),
				customerror.CommonErr,
			)
		}
		if ce.ErrCode != customerror.ItemNotFound {
			return "", err
		}

		if err := sus.sur.FindByID(ctx, decodedID, url); err != nil {
			return "", err
		}
		go func(eid string, u ShortenURL) {
			ttl := 24 * time.Hour
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			fun := func() error {
				return sus.suc.Set(ctx, eid, ttl, u)
			}

			services.WithErrorRetry(ctx, fun, 100*time.Millisecond)
		}(encodedID, *url)
	}

	if url.ExpiredAt != nil && now.After(*url.ExpiredAt) {
		return "", customerror.NewError(
			"expired url",
			errors.New("expired url"),
			customerror.InvalidAction,
		)
	}

	decryptedURL, err := sus.ue.DecryptURL(url.EncryptedLongUrl)
	if err != nil {
		return "", err
	}

	go func(u ShortenURL, did int64, dev string) {
		if u.UserID != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			fun := func() error {
				return sus.vrr.InsertNewRecord(ctx, *u.UserID, did, dev)
			}

			services.WithErrorRetry(ctx, fun, 100*time.Millisecond)
		}
	}(*url, decodedID, device)

	return decryptedURL, nil
}
