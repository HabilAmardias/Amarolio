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

func (sus *ShortenURLServiceImpl) FindOriginalURL(id string) (FindOriginalURL, error) {
	res, err := sus.callFindOriginalURL(id)
	if err != nil {
		return FindOriginalURL{}, err
	}
	return res.Data, nil
}

func (sus *ShortenURLServiceImpl) IsCustomURLAvailable(userID, customCode string) (bool, error) {
	return sus.callIsCustomURLAvailable(userID, customCode)
}

func (sus *ShortenURLServiceImpl) GetUserLinks(userID string, lastID *int64, limit int64) ([]URL, int64, int64, error) {
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

func (sus *ShortenURLServiceImpl) VisitOriginalURL(id string, device string) (string, error) {
	res, err := sus.callVisitOriginalURL(id, device)
	if err != nil {
		return "", err
	}

	return res.Data.URL, err
}

func (sus *ShortenURLServiceImpl) callIsCustomURLAvailable(userID string, customCode string) (bool, error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	body := IsCustomURLAvailableBody{
		CustomCode: customCode,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return false, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	_, err = services.Call[entity.PlainMessageResponse](sus.hs, sus.pr, "/api/v1/url/custom-code", http.MethodPost, http.StatusOK, b, nil, headers)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (sus *ShortenURLServiceImpl) callGetUserLinks(userID string, lastID *int64, limit int64) (*dto.ServerResponse[entity.Paginate[URL]], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	queries := map[string]string{
		"limit": strconv.FormatInt(limit, 10),
	}
	if lastID != nil {
		queries["last_id"] = strconv.FormatInt(*lastID, 10)
	}
	return services.Call[entity.Paginate[URL]](sus.hs, sus.pr, "/api/v1/me/url", http.MethodGet, http.StatusOK, nil, queries, headers)
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

func (sus *ShortenURLServiceImpl) callVisitOriginalURL(id string, device string) (*dto.ServerResponse[OriginalURL], error) {
	headers := map[string]string{
		"User-Agent": device,
	}
	return services.Call[OriginalURL](
		sus.hs,
		sus.pr,
		fmt.Sprintf("/api/v1/url/%s/redirect", id),
		http.MethodGet,
		http.StatusOK,
		nil,
		nil,
		headers,
	)
}

func (sus *ShortenURLServiceImpl) callFindOriginalURL(id string) (*dto.ServerResponse[FindOriginalURL], error) {
	return services.Call[FindOriginalURL](
		sus.hs,
		sus.pr,
		fmt.Sprintf("/api/v1/url/%s/metadata", id),
		http.MethodGet,
		http.StatusOK,
		nil,
		nil,
		nil,
	)
}
