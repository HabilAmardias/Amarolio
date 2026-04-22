package shortenurls

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenURLServiceImpl struct {
	hs string
	pr string
}

func NewShortenURLService(hs string, pr string) *ShortenURLServiceImpl {
	return &ShortenURLServiceImpl{hs, pr}
}

func (sus *ShortenURLServiceImpl) NewShortURL(userID *string, url string, duration *int) (string, error) {
	b := NewShortenURLBody{
		URL:      url,
		Duration: duration,
	}
	reqBody, err := json.Marshal(b)
	if err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	headers := map[string]string{}
	if userID != nil {
		headers[constants.X_USER_ID] = *userID
	}
	res, err := services.Call[NewShortenURL](
		sus.hs,
		sus.pr,
		"/api/v1/url",
		http.MethodPost,
		http.StatusCreated,
		reqBody,
		nil,
		headers,
	)
	if err != nil {
		return "", err
	}

	return res.Data.URL, nil
}

func (sus *ShortenURLServiceImpl) FindLongURL(id string, device string) (string, error) {
	headers := map[string]string{
		"User-Agent": device,
	}
	res, err := services.Call[FindLongURL](
		sus.hs,
		sus.pr,
		fmt.Sprintf("/api/v1/url/%s", id),
		http.MethodGet,
		http.StatusOK,
		nil,
		nil,
		headers,
	)
	if err != nil {
		return "", err
	}

	return res.Data.URL, err
}
