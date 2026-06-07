package url

import (
	"amary/src/customerror"
	"amary/src/services"
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

type URLEncryptorItf interface {
	DecryptURL(cipherURL string) (string, error)
	EncryptURL(plainURL string) (string, error)
}

type IDEncoderItf interface {
	Encode(id int64) string
	Decode(encodedID string) (int64, error)
	DecodeMultipleIDs(ids []string) ([]int64, error)
}

type URLCacheItf interface {
	Set(ctx context.Context, encodedID string, ttl time.Duration, url URL) error
	Get(ctx context.Context, encodedID string, url *URL) error
}

type VisitRecordRepoItf interface {
	InsertNewRecord(
		ctx context.Context,
		userID string,
		id int64,
		device string,
	) error
}

type URLRepoItf interface {
	InsertNewURL(
		ctx context.Context,
		userID *string,
		encryptedLongURL string,
		expiredAt *time.Time,
		shortenURL *URL,
	) error
	FindByID(ctx context.Context, id int64, url *URL) error
	FindUserLinks(ctx context.Context, userID string, lastID *int64, limit int64, links *[]URL) error
}

type URLServiceImpl struct {
	ue  URLEncryptorItf
	ide IDEncoderItf
	suc URLCacheItf
	sur URLRepoItf
	vrr VisitRecordRepoItf
}

func NewURLService(ue URLEncryptorItf, ide IDEncoderItf, suc URLCacheItf, sur URLRepoItf, vrr VisitRecordRepoItf) *URLServiceImpl {
	return &URLServiceImpl{ue, ide, suc, sur, vrr}
}

func (sus *URLServiceImpl) decryptAndFormatURL(ls []URL) ([]DecryptedURL, error) {
	decryptedLinks := []DecryptedURL{}
	for _, l := range ls {
		du, err := sus.ue.DecryptURL(l.EncryptedLongUrl)
		if err != nil {
			return nil, err
		}
		eid := sus.ide.Encode(l.ID)
		decryptedLinks = append(decryptedLinks, DecryptedURL{
			ID:        l.ID,
			UserID:    l.UserID,
			LongURL:   du,
			ShortURL:  fmt.Sprintf("%s/%s", os.Getenv("AMARY_REDIRECT_DOMAIN"), eid),
			CreatedAt: l.CreatedAt,
			ExpiredAt: l.ExpiredAt,
		})
	}
	return decryptedLinks, nil
}

func (sus *URLServiceImpl) GetUserLinks(ctx context.Context, userID string, lastID *int64, limit int64) ([]DecryptedURL, error) {
	links := new([]URL)
	if err := sus.sur.FindUserLinks(ctx, userID, lastID, limit, links); err != nil {
		return nil, err
	}
	// decrypt real url
	return sus.decryptAndFormatURL(*links)
}

func (sus *URLServiceImpl) NewShortURL(ctx context.Context, userID *string, longURL string, duration *int) (string, *time.Time, error) {
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

	url := new(URL)

	encryptedURL, err := sus.ue.EncryptURL(longURL)
	if err != nil {
		return "", nil, err
	}
	if err := sus.sur.InsertNewURL(ctx, userID, encryptedURL, eat, url); err != nil {
		return "", nil, err
	}

	encodedID := sus.ide.Encode(url.ID)

	go func(uid *string, eid string, u URL) {
		ttl := 24 * time.Hour
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// add to cache
		fun := func() error {
			return sus.suc.Set(ctx, eid, ttl, u)
		}

		services.WithErrorRetry(ctx, fun, 100*time.Millisecond)
	}(userID, encodedID, *url)

	return encodedID, eat, nil
}

func (sus *URLServiceImpl) FindLongURL(ctx context.Context, encodedID string, device string) (string, error) {
	url := new(URL)
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
		go func(eid string, u URL) {
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

	go func(u URL, did int64, dev string) {
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
