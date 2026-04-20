package shortenurl

import (
	"amary/src/customerror"
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
	Set(ctx context.Context, encodedID string, ttl time.Duration, url *ShortenURL) error
	Get(ctx context.Context, encodedID string, url *ShortenURL) error
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
	ttl := 24 * time.Hour
	now := time.Now()
	eatv := now.Add(ttl)
	eat := &eatv

	// user logged-in and no expiration time
	if userID != nil && duration == nil {
		ttl = 0
		eat = nil
	}

	// user logged-in and set expiration time
	if userID != nil && duration != nil {
		ttl = time.Duration(*duration) * 24 * time.Hour
		*eat = now.Add(ttl)
	}

	url := new(ShortenURL)

	encryptedURL := sus.ue.EncryptURL(longURL)
	if err := sus.sur.InsertNewURL(ctx, userID, encryptedURL, eat, url); err != nil {
		return "", err
	}

	encodedID := sus.ide.Encode(url.ID)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		attempt := 0
		maxRetry := 3
		if err := sus.suc.Set(ctx, encodedID, ttl, url); err != nil {
			for err != nil && attempt <= maxRetry {
				err = sus.suc.Set(ctx, encodedID, ttl, url)
				attempt += 1
			}
		}
	}()

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

	go func() {
		if url.UserID != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			attemptCount := 0
			maxRetry := 3
			if err := sus.vrr.InsertNewRecord(ctx, *url.UserID, decodedID, device); err != nil {
				for err != nil && attemptCount <= maxRetry {
					err = sus.vrr.InsertNewRecord(ctx, *url.UserID, decodedID, device)
					attemptCount += 1
				}
			}
		}
	}()

	return decryptedURL, nil
}
