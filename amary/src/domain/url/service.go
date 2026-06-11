package url

import (
	"amary/src/customerror"
	"amary/src/services"
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
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
	Set(ctx context.Context, shortCode string, ttl time.Duration, url URL) error
	Get(ctx context.Context, shortCode string, url *URL) error
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
	UpdateShortCode(ctx context.Context, id int64, shortCode string, url *URL) error
	FindByCode(ctx context.Context, shortCode string, url *URL) error
}

type TransactionManagerItf interface {
	WithTransaction(ctx context.Context, fun func(c context.Context) error) error
}

type URLServiceImpl struct {
	ue  URLEncryptorItf
	ide IDEncoderItf
	suc URLCacheItf
	sur URLRepoItf
	vrr VisitRecordRepoItf
	trm TransactionManagerItf
}

func NewURLService(ue URLEncryptorItf, ide IDEncoderItf, suc URLCacheItf, sur URLRepoItf, vrr VisitRecordRepoItf, trm TransactionManagerItf) *URLServiceImpl {
	return &URLServiceImpl{ue, ide, suc, sur, vrr, trm}
}

func (sus *URLServiceImpl) IsCustomURLAvailable(ctx context.Context, customCode string) (bool, error) {
	url := new(URL)
	// find on cache
	if err := sus.suc.Get(ctx, customCode, url); err != nil {
		// if failed, fallback to database
		if err := sus.sur.FindByCode(ctx, customCode, url); err != nil {
			var cErr *customerror.CustomError
			if !errors.As(err, &cErr) {
				return false, customerror.NewError(
					"something went wrong",
					errors.New("parse error failed"),
					customerror.CommonErr,
				)
			}
			// if database error, return the error
			if cErr.ErrCode != customerror.ItemNotFound {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (sus *URLServiceImpl) GetUserLinks(ctx context.Context, userID string, lastID *int64, limit int64) ([]DecryptedURL, error) {
	links := new([]URL)
	if err := sus.sur.FindUserLinks(ctx, userID, lastID, limit, links); err != nil {
		return nil, err
	}
	// decrypt real url
	return sus.decryptAndFormatURL(*links)
}

func (sus *URLServiceImpl) NewShortURL(ctx context.Context, userID *string, longURL string, duration *int, customCode *string) (string, *time.Time, error) {
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

	var shortCode string = sus.ide.Encode(url.ID)
	if customCode != nil && userID != nil {
		// basic validation
		if err := sus.validateCustomCode(*customCode); err != nil {
			return "", nil, err
		}

		// check dupe on cache, if exist, return error
		dupe := new(URL)
		if err := sus.suc.Get(ctx, *customCode, dupe); err == nil {
			return "", nil, customerror.NewError(
				"url already exist",
				errors.New("code already exist"),
				customerror.InvalidAction,
			)
		}

		// if not exist on cache, check dupe on db, if exist return error
		if err := sus.sur.FindByCode(ctx, *customCode, dupe); err != nil {
			var cErr *customerror.CustomError
			if !errors.As(err, &cErr) {
				return "", nil, customerror.NewError(
					"something went wrong",
					errors.New("failed to parse error"),
					customerror.CommonErr,
				)
			}

			if cErr.ErrCode != customerror.ItemNotFound {
				return "", nil, err
			}
		} else {
			return "", nil, customerror.NewError(
				"url already exist",
				errors.New("code already exist"),
				customerror.InvalidAction,
			)
		}

		shortCode = *customCode
	}

	if err := sus.trm.WithTransaction(ctx, func(c context.Context) error {
		if err := sus.sur.InsertNewURL(ctx, userID, encryptedURL, eat, url); err != nil {
			return err
		}
		return sus.sur.UpdateShortCode(ctx, url.ID, shortCode, url)
	}); err != nil {
		return "", nil, err
	}

	go func(uid *string, eid string, u URL) {
		ttl := 24 * time.Hour
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// add to cache
		fun := func() error {
			return sus.suc.Set(ctx, eid, ttl, u)
		}

		services.WithErrorRetry(ctx, fun, 100*time.Millisecond)

	}(userID, shortCode, *url)

	return shortCode, eat, nil
}

func (sus *URLServiceImpl) FindLongURL(ctx context.Context, shortCode string, device string) (string, error) {
	url := new(URL)
	now := time.Now()

	if err := sus.suc.Get(ctx, shortCode, url); err != nil {
		if err := sus.sur.FindByCode(ctx, shortCode, url); err != nil {
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
		}(shortCode, *url)
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

	go func(u URL, dev string) {
		if u.UserID != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			fun := func() error {
				return sus.vrr.InsertNewRecord(ctx, *u.UserID, u.ID, dev)
			}

			services.WithErrorRetry(ctx, fun, 100*time.Millisecond)
		}
	}(*url, device)

	return decryptedURL, nil
}

func (sus *URLServiceImpl) decryptAndFormatURL(ls []URL) ([]DecryptedURL, error) {
	decryptedLinks := []DecryptedURL{}
	for _, l := range ls {
		du, err := sus.ue.DecryptURL(l.EncryptedLongUrl)
		if err != nil {
			return nil, err
		}
		decryptedLinks = append(decryptedLinks, DecryptedURL{
			ID:        l.ID,
			UserID:    l.UserID,
			LongURL:   du,
			ShortURL:  fmt.Sprintf("%s/%s", os.Getenv("AMARY_REDIRECT_DOMAIN"), *l.ShortCode),
			CreatedAt: l.CreatedAt,
			ExpiredAt: l.ExpiredAt,
		})
	}
	return decryptedLinks, nil
}

func (sus *URLServiceImpl) validateCustomCode(code string) error {
	var alphaNumericRegex = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)

	if len(code) < 3 {
		return customerror.NewError(
			"custom url is too short",
			errors.New("custom code is too short"),
			customerror.InvalidAction,
		)
	}

	if !alphaNumericRegex.MatchString(code) {
		return customerror.NewError(
			"invalid code",
			errors.New("custom code does not comply to requirement"),
			customerror.InvalidAction,
		)
	}

	if ok := !strings.ContainsAny(code, "-_"); ok {
		return customerror.NewError(
			"invalid code",
			errors.New("collision risk"),
			customerror.InvalidAction,
		)
	}

	return nil
}
