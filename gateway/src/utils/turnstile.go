package utils

import (
	"amarolio-gateway/src/customerrors"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

type VerifyReq struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

type TurnstileUtil struct{}

func NewTurnstileUtil() *TurnstileUtil {
	return &TurnstileUtil{}
}

func (tu *TurnstileUtil) Validate(token string) error {
	url := "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	secret := os.Getenv("CF_TURNSTILE_SECRET")

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(http.MethodPost)
	req.Header.Set("Content-Type", "application/json")

	reqBody := VerifyReq{
		Secret:   secret,
		Response: token,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	req.SetBody(b)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	if res.StatusCode() != fasthttp.StatusOK {
		return customerrors.NewError(
			"validation failed",
			errors.New("failed to verify user"),
			customerrors.Unauthenticate,
		)
	}

	return nil
}
