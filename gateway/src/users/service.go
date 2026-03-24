package users

import (
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"encoding/json"
	"errors"
	"os"

	"github.com/valyala/fasthttp"
)

type UserServiceImpl struct{}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (us *UserServiceImpl) Login() (string, string, error) {

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(fasthttp.MethodGet)
	url := os.Getenv("AUTH_SERVICE_HOST") + ":" + os.Getenv("SERVICE_PORT") + "/api/v1/login"
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		resBody := new(dto.ServerResponse[dto.ErrorResponse])
		if err := json.Unmarshal(resp.Body(), resBody); err != nil {
			return "", "", customerrors.NewError(
				"something went wrong",
				err,
				customerrors.CommonErr,
			)
		}
		code := resp.StatusCode() * 100
		return "", "", customerrors.NewError(
			resBody.Data.Detail,
			errors.New(resBody.Data.Detail),
			code,
		)
	}
	resBody := new(dto.ServerResponse[LoginData])
	if err := json.Unmarshal(resp.Body(), resBody); err != nil {
		return "", "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	return resBody.Data.State, resBody.Data.URL, nil
}

func (us *UserServiceImpl) RefreshAuth(userID string) (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	url := os.Getenv("AUTH_SERVICE_HOST") + ":" + os.Getenv("SERVICE_PORT") + "/api/v1/refresh"

	req.SetRequestURI(url)
	req.Header.Set("x-user-id", userID)
	req.Header.SetMethod(fasthttp.MethodPost)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		resBody := new(dto.ServerResponse[dto.ErrorResponse])
		if err := json.Unmarshal(resp.Body(), resBody); err != nil {
			return "", customerrors.NewError(
				"something went wrong",
				err,
				customerrors.CommonErr,
			)
		}
		code := resp.StatusCode() * 100
		return "", customerrors.NewError(
			resBody.Data.Detail,
			errors.New(resBody.Data.Detail),
			code,
		)
	}

	resBody := new(dto.ServerResponse[AuthData])
	if err := json.Unmarshal(resp.Body(), resBody); err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	return resBody.Data.Token, nil
}
