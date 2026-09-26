package users

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"amarolio-auth/src/services"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type JWTUtilItf interface {
	GenerateJWT(id string, usedFor int, age time.Duration) (string, error)
}

type MailItf interface {
	SendEmail(email SendEmailParams) error
}

type UserCacheItf interface {
	FindCacheByID(ctx context.Context, userID string, user *User) error
	FindCacheByEmail(ctx context.Context, userEmail string, user *User) error
	SetCacheByID(ctx context.Context, user *User) error
	SetCacheByEmail(ctx context.Context, user *User) error
}

type HasherItf interface {
	HashPassword(password string) (string, error)
	ValidatePassword(encodedHash, password string) (bool, error)
}

type OTPGenItf interface {
	GenerateOTP() (string, error)
}

type Logger interface {
	Errorln(args ...interface{})
}

type UserServiceImpl struct {
	hu   HasherItf
	ou   OTPGenItf
	mu   MailItf
	ju   JWTUtilItf
	dbtx *db.DBHandle
	uc   UserCacheItf
	lg   Logger
}

func NewUserService(hu HasherItf, ou OTPGenItf, mu MailItf, ju JWTUtilItf, dbtx *db.DBHandle, uc UserCacheItf, lg Logger) *UserServiceImpl {
	return &UserServiceImpl{hu, ou, mu, ju, dbtx, uc, lg}
}

func (us *UserServiceImpl) ResendVerification(ctx context.Context, email string) error {
	user := new(User)
	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByEmail(ctx, email, user); err != nil {
		if err := ur.FindByEmail(ctx, email, user); err != nil {
			return err
		}

		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}

	if user.Verified {
		return customerrors.NewError(
			"user already verified",
			errors.New("user already verified"),
			customerrors.InvalidAction,
		)
	}

	token, err := us.generateVerificationToken()
	if err != nil {
		return err
	}
	eat := time.Now().Add(time.Hour)
	if err := ur.UpdateVerificationToken(ctx, user.ID, token, eat, user); err != nil {
		return err
	}

	go func() {
		url := fmt.Sprintf("%s/api/v1/verification?user_id=%s&token=%s", os.Getenv("SERVER_HOST"), user.ID, token)
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "User Verification | Amarolio",
			EmailBody: constants.BuildVerificationEmailBody(strings.Split(user.Email, "@")[0], url),
		}); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return nil
}

func (us *UserServiceImpl) VerifyUser(ctx context.Context, userID, token string) error {
	user := new(User)
	ur := NewUserRepository(us.dbtx)

	if err := us.uc.FindCacheByID(ctx, userID, user); err != nil {
		if err := ur.FindByID(ctx, userID, user); err != nil {
			return err
		}
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}

	if user.Verified {
		return customerrors.NewError(
			"user already verified",
			errors.New("user already verified"),
			customerrors.InvalidAction,
		)
	}
	if user.IsVerificationTokenExpired() {
		return customerrors.NewError(
			"verification token expired",
			errors.New("verification token expired"),
			customerrors.InvalidAction,
		)
	}
	if !user.IsVerificationTokenMatches(token) {
		return customerrors.NewError(
			"invalid verification token",
			errors.New("invalid verification token"),
			customerrors.InvalidAction,
		)
	}
	if err := ur.UpdateUserVerificationStatus(ctx, userID, true, user); err != nil {
		return err
	}

	return nil
}

func (us *UserServiceImpl) Register(ctx context.Context, email, password string) error {
	user := new(User)
	ur := NewUserRepository(us.dbtx)

	if err := us.uc.FindCacheByEmail(ctx, email, user); err == nil {
		return customerrors.NewError(
			"user already registered",
			errors.New("user already exist"),
			customerrors.InvalidAction,
		)
	}
	if err := ur.FindByEmail(ctx, email, user); err == nil {
		return customerrors.NewError(
			"user already registered",
			errors.New("user already exist"),
			customerrors.InvalidAction,
		)
	} else {
		var parsedErr *customerrors.CustomError
		if !errors.As(err, &parsedErr) {
			return customerrors.NewError(
				"something went wrong",
				errors.New("parse error failed"),
				customerrors.CommonErr,
			)
		}
		if parsedErr.ErrCode != customerrors.ItemNotFound {
			return err
		}
	}

	hashedPassword, err := us.hu.HashPassword(password)
	if err != nil {
		return err
	}

	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txUR := NewUserRepository(tx)
		if err := txUR.AddNewUser(ctx, email, hashedPassword, user); err != nil {
			return err
		}

		token, err := us.generateVerificationToken()
		if err != nil {
			return err
		}
		eat := time.Now().Add(time.Hour)

		return txUR.UpdateVerificationToken(ctx, user.ID, token, eat, user)
	}); err != nil {
		return err
	}

	go func() {
		url := fmt.Sprintf("%s/api/v1/verification?user_id=%s&token=%s", os.Getenv("SERVER_HOST"), user.ID, *user.VerificationToken)
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "User Verification | Amarolio",
			EmailBody: constants.BuildVerificationEmailBody(strings.Split(user.Email, "@")[0], url),
		}); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return nil
}

func (us *UserServiceImpl) GetProfile(ctx context.Context, userID string) (string, error) {
	user := new(User)
	if err := us.uc.FindCacheByID(ctx, userID, user); err != nil {
		// if cache miss, fetch from db
		ur := NewUserRepository(us.dbtx)
		if err := ur.FindByID(ctx, userID, user); err != nil {
			return "", err
		}
		// renew cache asynchronously
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}
	return strings.Split(user.Email, "@")[0], nil
}

func (us *UserServiceImpl) Login(ctx context.Context, userID string, otp string) (string, string, error) {
	user := new(User)
	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByID(ctx, userID, user); err != nil {
		if err := ur.FindByID(ctx, userID, user); err != nil {
			return "", "", err
		}
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}

	if !user.Verified {
		return "", "", customerrors.NewError(
			"user not verified",
			errors.New("user not verified"),
			customerrors.InvalidAction,
		)
	}

	if !user.IsOTPMatches(otp) {
		return "", "", customerrors.NewError(
			"Incorrect otp",
			errors.New("invalid otp"),
			customerrors.Unauthenticate,
		)
	}

	if user.IsOTPExpired() {
		return "", "", customerrors.NewError(
			"OTP has expired",
			errors.New("otp has expired"),
			customerrors.Unauthenticate,
		)
	}

	return us.generateAuthAndRefreshToken(userID)
}

func (us *UserServiceImpl) ResendOTP(ctx context.Context, userID string) (string, error) {
	user := new(User)
	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByID(ctx, userID, user); err != nil {
		if err := ur.FindByID(ctx, userID, user); err != nil {
			return "", err
		}
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}

	// Check if the OTP has expired, if not, return an error
	if !user.IsOTPExpired() {
		return "", customerrors.NewError(
			"OTP not expired yet",
			errors.New("otp not expired yet"),
			customerrors.Unauthenticate,
		)
	}

	otp, err := us.ou.GenerateOTP()
	if err != nil {
		return "", err
	}
	if err := ur.UpdateOTP(ctx, userID, otp, user); err != nil {
		return "", err
	}

	token, err := us.ju.GenerateJWT(user.ID, constants.ForOTP, constants.AUTH_AGE)
	if err != nil {
		return "", err
	}

	go func() {
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "One Time Password For Login",
			EmailBody: constants.BuildOTPEmailBody(strings.Split(user.Email, "@")[0], otp),
		}); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return token, nil
}

func (us *UserServiceImpl) PreLogin(ctx context.Context, email string, password string) (string, error) {
	user := new(User)

	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByEmail(ctx, email, user); err != nil {
		if err := ur.FindByEmail(ctx, email, user); err != nil {
			return "", err
		}
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}
	if !user.Verified {
		return "", customerrors.NewError(
			"user not verified",
			errors.New("user not verified"),
			customerrors.InvalidAction,
		)
	}

	match, err := us.hu.ValidatePassword(user.Password, password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", customerrors.NewError(
			"invalid credentials",
			errors.New("invalid password"),
			customerrors.InvalidAction,
		)
	}

	otp, err := us.ou.GenerateOTP()
	if err != nil {
		return "", err
	}
	if err := ur.UpdateOTP(ctx, user.ID, otp, user); err != nil {
		return "", err
	}

	token, err := us.ju.GenerateJWT(user.ID, constants.ForOTP, constants.AUTH_AGE)
	if err != nil {
		return "", err
	}

	go func() {
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "One Time Password For Login",
			EmailBody: constants.BuildOTPEmailBody(strings.Split(user.Email, "@")[0], otp),
		}); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return token, nil
}

func (us *UserServiceImpl) RefreshAuth(ctx context.Context, userID string) (string, error) {
	user := new(User)

	// try fetch from cache
	if err := us.uc.FindCacheByID(ctx, userID, user); err != nil {
		// if cache miss, fetch from db
		ur := NewUserRepository(us.dbtx)
		if err := ur.FindByID(ctx, userID, user); err != nil {
			return "", err
		}
		// renew cache asynchronously
		go func() {
			if err := us.renewCache(user); err != nil {
				us.lg.Errorln(err.Error())
			}
		}()
	}
	if !user.Verified {
		return "", customerrors.NewError(
			"user not verified",
			nil,
			customerrors.CommonErr,
		)
	}
	token, err := us.ju.GenerateJWT(userID, constants.ForAuth, constants.AUTH_AGE)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserServiceImpl) generateAuthAndRefreshToken(userID string) (string, string, error) {
	authToken, err := us.ju.GenerateJWT(userID, constants.ForAuth, constants.AUTH_AGE)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := us.ju.GenerateJWT(userID, constants.ForRefresh, constants.REFRESH_AGE)
	if err != nil {
		return "", "", err
	}

	return authToken, refreshToken, nil
}

func (us *UserServiceImpl) renewCache(user *User) error {
	newCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := us.uc.SetCacheByID(newCtx, user); err != nil {
		return err
	}
	return us.uc.SetCacheByEmail(newCtx, user)
}

func (us *UserServiceImpl) generateVerificationToken() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	return hex.EncodeToString(b), nil
}
