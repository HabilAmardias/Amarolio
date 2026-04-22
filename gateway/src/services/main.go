package services

import (
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

func Call[T any](host string, port string, path string, method string, successStatus int, body []byte, queries map[string]string, headers map[string]string) (*dto.ServerResponse[T], error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	queryString := []string{}
	for k, v := range queries {
		queryString = append(queryString, fmt.Sprintf("%s=%s", k, v))
	}
	url := host + ":" + port + path
	if len(queryString) > 0 {
		url += fmt.Sprintf("?%s", strings.Join(queryString, "&"))
	}
	req.SetRequestURI(url)
	if body != nil {
		req.SetBody(body)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	if res.StatusCode() != successStatus {
		resBody := new(dto.ServerResponse[dto.ErrorResponse])
		if err := json.Unmarshal(res.Body(), resBody); err != nil {
			return nil, customerrors.NewError(
				"something went wrong",
				err,
				customerrors.CommonErr,
			)
		}
		code := res.StatusCode() * 100
		return nil, customerrors.NewError(
			resBody.Data.Detail,
			errors.New(resBody.Data.Detail),
			code,
		)
	}
	resBody := new(dto.ServerResponse[T])
	if err := json.Unmarshal(res.Body(), resBody); err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return resBody, nil
}
