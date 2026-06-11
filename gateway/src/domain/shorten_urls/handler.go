package shortenurls

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type ShortenURLServiceItf interface {
	FindLongURL(id string, device string) (string, error)
	NewShortURL(userID *string, url string, duration *int, customCode *string) (NewShortenURL, error)
	GetUserLinks(userID string, lastID *int64, limit int64) ([]UserLink, int64, int64, error)
	IsCustomURLAvailable(userID, customCode string) (bool, error)
}

type ShortenURLHandlerImpl struct {
	sus ShortenURLServiceItf
}

func NewShortenURLHandler(sus ShortenURLServiceItf) *ShortenURLHandlerImpl {
	return &ShortenURLHandlerImpl{sus}
}

func (suh *ShortenURLHandlerImpl) IsCustomURLAvailable(ctx fiber.Ctx) error {
	req := new(IsCustomURLAvailableReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	available, err := suh.sus.IsCustomURLAvailable(claim.Subject, req.CustomCode)
	if err != nil {
		return err
	}
	if !available {
		return customerrors.NewError(
			"url is not available",
			errors.New("url is not available"),
			customerrors.InvalidAction,
		)
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.PlainMessageRes]{
		Success: true,
		Data: dto.PlainMessageRes{
			Message: "URL available",
		},
	})
}

func (suh *ShortenURLHandlerImpl) GetUserLinks(ctx fiber.Ctx) error {
	req := new(GetUserLinksReq)
	if err := ctx.Bind().Query(req); err != nil {
		return customerrors.NewError(
			"invalid input",
			err,
			customerrors.InvalidAction,
		)
	}
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	var reqLimit int64 = constants.DEFAULT_LIMIT_PAGINATION
	if req.Limit != nil {
		reqLimit = *req.Limit
	}
	entries, lastID, limit, err := suh.sus.GetUserLinks(claim.Subject, req.LastID, reqLimit)
	if err != nil {
		return err
	}
	res := []GetUserLinkRes{}
	for _, e := range entries {
		res = append(res, GetUserLinkRes(e))
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.PaginateRes[GetUserLinkRes]]{
		Success: true,
		Data: dto.PaginateRes[GetUserLinkRes]{
			Entries: res,
			PageInfo: struct {
				LastID   *int64 "json:\"last_id,omitempty\""
				Page     *int64 "json:\"page,omitempty\""
				Limit    int64  "json:\"limit\""
				FilterBy []struct {
					Name  string "json:\"name\""
					Value any    "json:\"value\""
				} "json:\"filter_by,omitempty\""
				SortBy []struct {
					Name   string "json:\"name\""
					Ascend bool   "json:\"ascend\""
				} "json:\"sort_by,omitempty\""
			}{
				LastID: &lastID,
				Limit:  limit,
			},
		},
	})
}

func (suh *ShortenURLHandlerImpl) RedirectToURL(ctx fiber.Ctx) error {
	device := ctx.Request().Header.UserAgent()
	id := ctx.Params("id")

	url, err := suh.sus.FindLongURL(id, string(device))
	if err != nil {
		return err
	}

	return ctx.Redirect().Status(http.StatusSeeOther).To(url)
}

func (suh *ShortenURLHandlerImpl) NewShortURL(ctx fiber.Ctx) error {
	var userID *string = nil
	req := new(NewShortenURLReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return customerrors.NewError(
			"invalid input",
			err,
			customerrors.InvalidAction,
		)
	}
	claim, _ := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if claim != nil {
		userID = &claim.Subject
	}

	res, err := suh.sus.NewShortURL(userID, req.URL, req.Duration, req.CustomCode)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(dto.ServerResponse[NewShortenURLRes]{
		Success: true,
		Data: NewShortenURLRes{
			URL:         res.URL,
			OriginalURL: res.OriginalURL,
			ExpiredAt:   res.ExpiredAt,
		},
	})
}
