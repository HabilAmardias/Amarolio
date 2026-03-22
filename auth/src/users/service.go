package users

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type OauthUtilItf interface {
	GetAuthCodeURL(state string) string
	Exchange(ctx context.Context, code string) (string, error)
	GetUrlAPI() string
}

type JWTUtilItf interface {
	GenerateJWT(id string, usedFor int, age time.Duration) (string, error)
}

type UserServiceImpl struct {
	ou   OauthUtilItf
	ju   JWTUtilItf
	dbtx *db.DBHandle
}

func NewUserService(ou OauthUtilItf, ju JWTUtilItf, dbtx *db.DBHandle) *UserServiceImpl {
	return &UserServiceImpl{ou, ju, dbtx}
}

func (us *UserServiceImpl) Login() (string, string) {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.RawURLEncoding.EncodeToString(b)

	url := us.ou.GetAuthCodeURL(state)

	return url, state
}

func (us *UserServiceImpl) RefreshAuth(ctx context.Context, userID string) (string, error) {
	user := new(User)
	ur := NewUserRepository(us.dbtx)

	if err := ur.FindByID(ctx, userID, user); err != nil {
		return "", err
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

func (us *UserServiceImpl) LoginCallback(ctx context.Context, code string) (string, string, error) {
	accessToken, err := us.ou.Exchange(ctx, code)
	if err != nil {
		return "", "", err
	}
	response, err := http.Get(us.ou.GetUrlAPI() + accessToken)
	if err != nil {
		return "", "", err
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	var userInfo map[string]interface{}
	if err := json.Unmarshal(contents, &userInfo); err != nil {
		return "", "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	email, ok := userInfo["email"].(string)
	if !ok {
		return "", "", customerrors.NewError(
			"something went wrong",
			errors.New("email not found"),
			customerrors.CommonErr,
		)
	}

	user := new(User)
	ur := NewUserRepository(us.dbtx)

	if err := ur.FindByEmail(ctx, email, user); err != nil {
		var parsedErr *customerrors.CustomError
		if !errors.As(err, &parsedErr) {
			return "", "", customerrors.NewError(
				"something went wrong",
				errors.New("fail parse error"),
				customerrors.CommonErr,
			)
		}
		if parsedErr.ErrCode != customerrors.ItemNotFound {
			return "", "", err
		}
		if err := ur.AddNewUser(ctx, email, user); err != nil {
			return "", "", err
		}
	}

	return us.generateAuthAndRefreshToken(user.ID)
}
