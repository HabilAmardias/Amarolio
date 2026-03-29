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

func (us *UserServiceImpl) callLogin() (*dto.ServerResponse[Login], error) {
	return services.Call[Login](us.hs, us.pr, "/api/v1/login", fasthttp.MethodPost, fasthttp.StatusOK, nil, nil, nil)
}

func (us *UserServiceImpl) callRefreshAuth() (*dto.ServerResponse[RefreshAuth], error) {
	return services.Call[RefreshAuth](us.hs, us.pr, "/api/v1/refresh", fasthttp.MethodPost, fasthttp.StatusOK, nil, nil, nil)
}

func (us *UserServiceImpl) callLoginCallback(code string, state string) (*dto.ServerResponse[LoginCallback], error) {
	queries := map[string]string{
		"code":  code,
		"state": state,
	}
	b := new(LoginCallbackBody)
	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[LoginCallback](us.hs, us.pr, "/api/v1/login/callback", fasthttp.MethodPost, fasthttp.StatusOK, reqBody, queries, nil)
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
