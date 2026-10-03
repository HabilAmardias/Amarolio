package users

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	otpchallenges "amarolio-auth/src/domain/otp_challenges"
	resetpasswordchallenges "amarolio-auth/src/domain/reset_password_challenges"
	verifyuserchallenges "amarolio-auth/src/domain/verify_user_challenges"
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
	SetCacheByID(ctx context.Context, age time.Duration, user *User) error
	SetCacheByEmail(ctx context.Context, age time.Duration, user *User) error
	GetLoginAttempts(ctx context.Context, email string) (int, error)
	IncreaseLoginAttempts(ctx context.Context, email string, max int, lock time.Duration) (int, error)
	ResetLoginAttempts(ctx context.Context, email string) error
	LoginLockRemaining(ctx context.Context, email string) (time.Duration, error)
}

type OTPGenItf interface {
	GenerateOTP() (string, error)
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

func lockoutError(ttl time.Duration) error {
	minutes := int(ttl.Round(time.Minute).Minutes())
	if minutes < 1 {
		minutes = 1
	}
	return customerrors.NewError(
		fmt.Sprintf("Too many failed attempts. Try again in %d minutes.", minutes),
		errors.New("too many failed login attempts"),
		customerrors.TooManyAttempts,
	)
}

// loginLockError returns a 429 error when the account has reached the failed
// attempt threshold. Infrastructure errors fail open so a cache outage cannot
// block every login.
func (us *UserServiceImpl) loginLockError(ctx context.Context, email string) error {
	attempts, err := us.uc.GetLoginAttempts(ctx, email)
	if err != nil {
		us.lg.Errorln("login lock check:", err.Error())
		return nil
	}
	if attempts < constants.MaxLoginAttempts {
		return nil
	}
	ttl, err := us.uc.LoginLockRemaining(ctx, email)
	if err != nil {
		us.lg.Errorln("login lock ttl:", err.Error())
	}
	return lockoutError(ttl)
}

// recordFailedLogin counts a failed attempt and returns the lockout error once
// the threshold is reached, otherwise the provided fallback error.
func (us *UserServiceImpl) recordFailedLogin(ctx context.Context, email string, fallback error) error {
	count, err := us.uc.IncreaseLoginAttempts(ctx, email, constants.MaxLoginAttempts, constants.LoginLockDuration)
	if err != nil {
		us.lg.Errorln("record login attempt:", err.Error())
		return fallback
	}
	if count >= constants.MaxLoginAttempts {
		ttl, ttlErr := us.uc.LoginLockRemaining(ctx, email)
		if ttlErr != nil {
			us.lg.Errorln("login lock ttl:", ttlErr.Error())
		}
		return lockoutError(ttl)
	}
	return fallback
}

func (us *UserServiceImpl) ResetPassword(ctx context.Context, token string, newPassword string) error {
	user := new(User)
	challenge := new(resetpasswordchallenges.ResetPasswordChallenge)

	rur := resetpasswordchallenges.NewResetPasswordChallengeRepository(us.dbtx)
	ur := NewUserRepository(us.dbtx)

	if err := rur.FindByID(ctx, token, challenge); err != nil {
		return err
	}

	if err := us.uc.FindCacheByID(ctx, challenge.UserID, user); err != nil {
		if err := ur.FindByID(ctx, challenge.UserID, user); err != nil {
			return err
		}
	}

	if challenge.IsExpired() {
		return customerrors.NewError(
			"Unauthorized",
			errors.New("reset password token expired"),
			customerrors.Unauthenticate,
		)
	}

	match, err := us.hu.Validate(user.Password, newPassword)
	if err != nil {
		return err
	}
	if match {
		return customerrors.NewError(
			"new password cannot be the same as the old password",
			errors.New("new password cannot be the same as the old password"),
			customerrors.InvalidAction,
		)
	}
	hashedPass, err := us.hu.Hash(newPassword)
	if err != nil {
		return err
	}

	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txUR := NewUserRepository(tx)
		txRUR := resetpasswordchallenges.NewResetPasswordChallengeRepository(tx)

		if err := txUR.UpdateUserPassword(ctx, user.ID, hashedPass, user); err != nil {
			return err
		}
		return txRUR.DeleteAllByUserID(ctx, user.ID)
	}); err != nil {
		return err
	}

	if err := us.uc.ResetLoginAttempts(ctx, user.Email); err != nil {
		us.lg.Errorln("reset login attempts:", err.Error())
	}

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err)
		}
	}()
	return nil
}

func (us *UserServiceImpl) SendResetPasswordEmail(ctx context.Context, email string) error {
	user := new(User)
	ur := NewUserRepository(us.dbtx)

	if err := us.uc.FindCacheByEmail(ctx, email, user); err != nil {
		if err := ur.FindByEmail(ctx, email, user); err != nil {
			// return nil even when the email does not exist, to prevent enumeration attack
			return nil
		}
	}

	challenge := new(resetpasswordchallenges.ResetPasswordChallenge)
	services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txRUR := resetpasswordchallenges.NewResetPasswordChallengeRepository(tx)
		eat := time.Now().Add(time.Hour)
		if err := txRUR.DeleteAllByUserID(ctx, user.ID); err != nil {
			return err
		}
		token, err := us.generateRandomToken()
		if err != nil {
			return err
		}
		return txRUR.CreateNewResetPasswordChallenge(ctx, token, user.ID, eat, challenge)
	})

	go func() {
		url := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("AUTH_CLIENT_URL"), challenge.ID)
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "Reset Password | Amarolio",
			EmailBody: constants.BuildResetPasswordEmailBody(strings.Split(user.Email, "@")[0], url),
		}); err != nil {
			us.lg.Errorln(err)
			return
		}
		us.lg.Infoln("email sent")
	}()

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err)
		}
	}()
	return nil
}

func (us *UserServiceImpl) ResendVerification(ctx context.Context, email string) error {
	user := new(User)
	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByEmail(ctx, email, user); err != nil {
		if err := ur.FindByEmail(ctx, email, user); err != nil {
			var parsedErr *customerrors.CustomError
			if !errors.As(err, &parsedErr) {
				return customerrors.NewError(
					"something went wrong",
					errors.New("parse error failed"),
					customerrors.CommonErr,
				)
			}
			if parsedErr.ErrCode == customerrors.ItemNotFound {
				return nil
			}
			return err
		}
	}

	if user.Verified {
		return customerrors.NewError(
			"user already verified",
			errors.New("user already verified"),
			customerrors.InvalidAction,
		)
	}

	token, err := us.generateRandomToken()
	if err != nil {
		return err
	}
	eat := time.Now().Add(time.Hour)

	challenge := new(verifyuserchallenges.VerifyUserChallenge)
	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txvur := verifyuserchallenges.NewVerifyUserChallengeRepository(tx)
		if err := txvur.DeleteAllByUserID(ctx, user.ID); err != nil {
			return err
		}
		return txvur.CreateNewVerifyUserChallenge(ctx, token, user.ID, eat, challenge)
	}); err != nil {
		return err
	}

	go func() {
		url := fmt.Sprintf("%s/verify?token=%s", os.Getenv("AUTH_CLIENT_URL"), challenge.ID)
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "User Verification | Amarolio",
			EmailBody: constants.BuildVerificationEmailBody(strings.Split(user.Email, "@")[0], url),
		}); err != nil {
			us.lg.Errorln(err.Error())
			return
		}
		us.lg.Infoln("email sent")
	}()

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return nil
}

func (us *UserServiceImpl) VerifyUser(ctx context.Context, token string) error {
	user := new(User)
	challenge := new(verifyuserchallenges.VerifyUserChallenge)

	ur := NewUserRepository(us.dbtx)
	vur := verifyuserchallenges.NewVerifyUserChallengeRepository(us.dbtx)

	if err := vur.FindByID(ctx, token, challenge); err != nil {
		return err
	}

	if err := us.uc.FindCacheByID(ctx, challenge.UserID, user); err != nil {
		if err := ur.FindByID(ctx, challenge.UserID, user); err != nil {
			return err
		}
	}

	if user.Verified {
		return customerrors.NewError(
			"user already verified",
			errors.New("user already verified"),
			customerrors.InvalidAction,
		)
	}

	if challenge.IsExpired() {
		return customerrors.NewError(
			"verification token expired",
			errors.New("verification token expired"),
			customerrors.InvalidAction,
		)
	}

	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txur := NewUserRepository(tx)
		txvur := verifyuserchallenges.NewVerifyUserChallengeRepository(tx)
		if err := txur.UpdateUserVerificationStatus(ctx, challenge.UserID, true, user); err != nil {
			return err
		}
		return txvur.DeleteAllByUserID(ctx, user.ID)
	}); err != nil {
		return err
	}

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

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

	hashedPassword, err := us.hu.Hash(password)
	if err != nil {
		return err
	}
	challenge := new(verifyuserchallenges.VerifyUserChallenge)
	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txUR := NewUserRepository(tx)
		txVUR := verifyuserchallenges.NewVerifyUserChallengeRepository(tx)
		if err := txUR.AddNewUser(ctx, email, hashedPassword, user); err != nil {
			return err
		}

		token, err := us.generateRandomToken()
		if err != nil {
			return err
		}
		eat := time.Now().Add(time.Hour)

		return txVUR.CreateNewVerifyUserChallenge(ctx, token, user.ID, eat, challenge)
	}); err != nil {
		return err
	}

	go func() {
		url := fmt.Sprintf("%s/verify?token=%s", os.Getenv("AUTH_CLIENT_URL"), challenge.ID)
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "User Verification | Amarolio",
			EmailBody: constants.BuildVerificationEmailBody(strings.Split(user.Email, "@")[0], url),
		}); err != nil {
			us.lg.Errorln(err.Error())
			return
		}
		us.lg.Infoln("email sent")
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

func (us *UserServiceImpl) Login(ctx context.Context, challengeID string, otp string) (string, string, error) {
	user := new(User)
	challenge := new(otpchallenges.OTPChallenge)

	ur := NewUserRepository(us.dbtx)
	ocr := otpchallenges.NewOTPChallengeRepository(us.dbtx)

	if err := ocr.FindByID(ctx, challengeID, challenge); err != nil {
		return "", "", err
	}

	if err := us.uc.FindCacheByID(ctx, challenge.UserID, user); err != nil {
		if err := ur.FindByID(ctx, challenge.UserID, user); err != nil {
			return "", "", err
		}
	}

	if err := us.loginLockError(ctx, user.Email); err != nil {
		return "", "", err
	}

	if !user.Verified {
		return "", "", customerrors.NewError(
			"user not verified",
			errors.New("user not verified"),
			customerrors.InvalidAction,
		)
	}

	// An expired OTP is not a wrong credential, so it is checked first and not
	// counted toward the lockout.
	if challenge.IsExpired() {
		return "", "", customerrors.NewError(
			"OTP has expired",
			errors.New("otp has expired"),
			customerrors.Unauthenticate,
		)
	}
	match, err := challenge.VerifyOTP(otp, us.hu)
	if err != nil {
		return "", "", err
	}
	if !match {
		return "", "", us.recordFailedLogin(ctx, user.Email, customerrors.NewError(
			"Incorrect otp",
			errors.New("invalid otp"),
			customerrors.Unauthenticate,
		))
	}

	// ---------------- Clean Up ----------------
	// invalidate all OTP Challenge for the user
	if err := ocr.DeleteAllByUserID(ctx, user.ID); err != nil {
		return "", "", err
	}
	// Reset login attempts for the user
	if err := us.uc.ResetLoginAttempts(ctx, user.Email); err != nil {
		us.lg.Errorln("reset login attempts:", err.Error())
	}
	// renew cache for the user
	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return us.generateAuthAndRefreshToken(user.ID)
}

func (us *UserServiceImpl) ResendOTP(ctx context.Context, challengeID string) (string, error) {
	user := new(User)
	challenge := new(otpchallenges.OTPChallenge)

	ur := NewUserRepository(us.dbtx)
	ocr := otpchallenges.NewOTPChallengeRepository(us.dbtx)

	if err := ocr.FindByID(ctx, challengeID, challenge); err != nil {
		return "", err
	}

	if err := us.uc.FindCacheByID(ctx, challenge.UserID, user); err != nil {
		if err := ur.FindByID(ctx, challenge.UserID, user); err != nil {
			return "", err
		}
	}

	if !user.Verified {
		return "", customerrors.NewError(
			"User not verified",
			errors.New("user not verified"),
			customerrors.Unauthenticate,
		)
	}

	if err := us.loginLockError(ctx, user.Email); err != nil {
		return "", err
	}

	// Check if the OTP has expired, if not, return an error
	if !challenge.IsExpired() {
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

	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		txocr := otpchallenges.NewOTPChallengeRepository(tx)
		// invalidate all existing challenges for this user
		if err := txocr.DeleteAllByUserID(ctx, user.ID); err != nil {
			return err
		}

		challengeID, err := us.generateRandomToken()
		if err != nil {
			return err
		}
		otpHash, err := us.hu.Hash(otp)
		if err != nil {
			return err
		}
		return txocr.CreateNewOTPChallenge(ctx, challengeID, user.ID, otpHash, time.Now().Add(time.Minute), challenge)
	}); err != nil {
		return "", err
	}

	go func() {
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "Your Amarolio sign-in code",
			EmailBody: constants.BuildOTPEmailBody(strings.Split(user.Email, "@")[0], otp),
		}); err != nil {
			us.lg.Errorln(err.Error())
			return
		}
		us.lg.Errorln("email sent")
	}()

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return challenge.ID, nil
}

func (us *UserServiceImpl) PreLogin(ctx context.Context, email string, password string) (string, error) {
	if err := us.loginLockError(ctx, email); err != nil {
		return "", err
	}

	user := new(User)

	ur := NewUserRepository(us.dbtx)
	if err := us.uc.FindCacheByEmail(ctx, email, user); err != nil {
		if err := ur.FindByEmail(ctx, email, user); err != nil {
			var parsedErr *customerrors.CustomError
			if !errors.As(err, &parsedErr) {
				return "", customerrors.NewError(
					"something went wrong",
					errors.New("parse error fail"),
					customerrors.CommonErr,
				)
			}
			if parsedErr.ErrCode == customerrors.ItemNotFound {
				return "", us.recordFailedLogin(ctx, email, customerrors.NewError(
					"invalid credentials",
					err,
					customerrors.InvalidAction,
				))
			}
			return "", err
		}
	}
	if !user.Verified {
		return "", customerrors.NewError(
			"user not verified",
			errors.New("user not verified"),
			customerrors.InvalidAction,
		)
	}

	match, err := us.hu.Validate(user.Password, password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", us.recordFailedLogin(ctx, email, customerrors.NewError(
			"invalid credentials",
			errors.New("invalid password"),
			customerrors.InvalidAction,
		))
	}

	otp, err := us.ou.GenerateOTP()
	if err != nil {
		return "", err
	}

	challenge := new(otpchallenges.OTPChallenge)
	if err := services.WithTransaction(us.dbtx, func(tx db.DBTXItf) error {
		ocr := otpchallenges.NewOTPChallengeRepository(tx)
		// invalidate all existing challenges for this user
		if err := ocr.DeleteAllByUserID(ctx, user.ID); err != nil {
			return err
		}
		challengeID, err := us.generateRandomToken()
		if err != nil {
			return err
		}
		otpHash, err := us.hu.Hash(otp)
		if err != nil {
			return err
		}
		return ocr.CreateNewOTPChallenge(ctx, challengeID, user.ID, otpHash, time.Now().Add(time.Minute), challenge)
	}); err != nil {
		return "", err
	}

	go func() {
		if err := us.mu.SendEmail(SendEmailParams{
			Receiver:  user.Email,
			Subject:   "Your Amarolio sign-in code",
			EmailBody: constants.BuildOTPEmailBody(strings.Split(user.Email, "@")[0], otp),
		}); err != nil {
			us.lg.Errorln(err.Error())
			return
		}
		us.lg.Infoln("email sent")
	}()

	go func() {
		if err := us.renewCache(user); err != nil {
			us.lg.Errorln(err.Error())
		}
	}()

	return challenge.ID, nil
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
	cacheAge := 15 * time.Minute
	if err := us.uc.SetCacheByID(newCtx, cacheAge, user); err != nil {
		return err
	}
	return us.uc.SetCacheByEmail(newCtx, cacheAge, user)
}

func (us *UserServiceImpl) generateRandomToken() (string, error) {
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
