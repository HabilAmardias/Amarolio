package shortenurls

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/entity"
	"amarolio-gateway/src/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ShortenURLServiceImpl struct {
	hs string
	pr string
}

func NewShortenURLService(hs string, pr string) *ShortenURLServiceImpl {
	return &ShortenURLServiceImpl{hs, pr}
}

func (sus *ShortenURLServiceImpl) GetUserLinks(userID string, lastID *int64, limit int64) ([]UserLink, int64, int64, error) {
	res, err := sus.callGetUserLinks(userID, lastID, limit)
	if err != nil {
		return nil, 0, 0, err
	}
	return res.Data.Entries, *res.Data.PageInfo.LastID, res.Data.PageInfo.Limit, nil
}

func (sus *ShortenURLServiceImpl) NewShortURL(userID *string, url string, duration *int, customCode *string) (NewShortenURL, error) {
	res, err := sus.callNewShortURL(userID, url, duration, customCode)
	if err != nil {
		return NewShortenURL{}, err
	}

	return res.Data, nil
}

func (sus *ShortenURLServiceImpl) FindLongURL(id string, device string) (string, error) {
	res, err := sus.callFindLongURL(id, device)
	if err != nil {
		return "", err
	}

	return res.Data.URL, err
}

func (sus *ShortenURLServiceImpl) callGetUserLinks(userID string, lastID *int64, limit int64) (*dto.ServerResponse[entity.Paginate[UserLink]], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	queries := map[string]string{
		"limit": strconv.FormatInt(limit, 10),
	}
	if lastID != nil {
		queries["last_id"] = strconv.FormatInt(*lastID, 10)
	}
	return services.Call[entity.Paginate[UserLink]](sus.hs, sus.pr, "/api/v1/me/url", http.MethodGet, http.StatusOK, nil, queries, headers)
}

func (sus *ShortenURLServiceImpl) callNewShortURL(userID *string, url string, duration *int, customCode *string) (*dto.ServerResponse[NewShortenURL], error) {
	b := NewShortenURLBody{
		URL:        url,
		Duration:   duration,
		CustomCode: customCode,
	}
	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	headers := map[string]string{}
	if userID != nil {
		headers[constants.X_USER_ID] = *userID
	}
	return services.Call[NewShortenURL](
		sus.hs,
		sus.pr,
		"/api/v1/url",
		http.MethodPost,
		http.StatusCreated,
		reqBody,
		nil,
		headers,
	)
}

func (sus *ShortenURLServiceImpl) callFindLongURL(id string, device string) (*dto.ServerResponse[FindLongURL], error) {
	headers := map[string]string{
		"User-Agent": device,
	}
	return services.Call[FindLongURL](
		sus.hs,
		sus.pr,
		fmt.Sprintf("/api/v1/url/%s", id),
		http.MethodGet,
		http.StatusOK,
		nil,
		nil,
		headers,
	)
}
