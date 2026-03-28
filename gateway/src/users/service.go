package users

import (
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/services"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type UserServiceImpl struct {
	hs string
	pr string
}

func NewUserService(hs string, pr string) *UserServiceImpl {
	return &UserServiceImpl{hs, pr}
}

func (us *UserServiceImpl) callLogin() (*dto.ServerResponse[LoginData], error) {
	return services.Call[LoginData](us.hs, us.pr, "/api/v1/login", fasthttp.MethodPost, fasthttp.StatusOK, nil, nil)
}

func (us *UserServiceImpl) callRefreshAuth() (*dto.ServerResponse[AuthData], error) {
	return services.Call[AuthData](us.hs, us.pr, "/api/v1/refresh", fasthttp.MethodPost, fasthttp.StatusOK, nil, nil)
}

func (us *UserServiceImpl) callLoginCallback(code string, state string) (*dto.ServerResponse[LoginCallbackData], error) {
	queries := map[string]string{
		"code":  code,
		"state": state,
	}
	b := new(AuthRefreshData)
	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[LoginCallbackData](us.hs, us.pr, "/api/v1/login/callback", fasthttp.MethodPost, fasthttp.StatusOK, reqBody, queries)
}

func (us *UserServiceImpl) LoginCallback(code string, state string) (string, string, error) {
	res, err := us.callLoginCallback(code, state)
	if err != nil {
		return "", "", err
	}
	return res.Data.AuthToken, res.Data.RefreshToken, nil
}

func (us *UserServiceImpl) Login() (string, string, error) {
	res, err := us.callLogin()
	if err != nil {
		return "", "", err
	}
	return res.Data.State, res.Data.URL, nil
}

func (us *UserServiceImpl) RefreshAuth(userID string) (string, error) {
	res, err := us.callRefreshAuth()
	if err != nil {
		return "", err
	}

	return res.Data.Token, nil
}
