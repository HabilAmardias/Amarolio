package url

import (
	"amary/src/constant"
	"amary/src/customerror"
	"amary/src/dto"
	"amary/src/handlers"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type URLServiceItf interface {
	NewShortURL(ctx context.Context, userID *string, longURL string, duration *int, customCode *string) (string, *time.Time, error)
	VisitOriginalURL(ctx context.Context, encodedID string, device string) (string, error)
	GetUserLinks(ctx context.Context, userID string, lastID *int64, limit int64) ([]DecryptedURL, error)
	IsCustomURLAvailable(ctx context.Context, customCode string) (bool, error)
	FindOriginalURL(ctx context.Context, shortCode string) (DecryptedURL, error)
}

type URLHandlerImpl struct {
	sus URLServiceItf
}

func NewURLHandler(sus URLServiceItf) *URLHandlerImpl {
	return &URLHandlerImpl{sus}
}

func (suh *URLHandlerImpl) FindOriginalURL(ctx *gin.Context) {
	id := ctx.Param("id")
	url, err := suh.sus.FindOriginalURL(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.ServerResponse[FindOriginalUrlRes]{
		Success: true,
		Data: FindOriginalUrlRes{
			URL: URLRes(url),
		},
	})
}

func (suh *URLHandlerImpl) IsCustomURLExist(ctx *gin.Context) {
	uid := handlers.GetAuthenticationPayload(ctx)
	if len(uid) == 0 {
		ctx.Error(customerror.NewError(
			"unauthorized",
			errors.New("user id does not provided"),
			customerror.Unauthenticate,
		))
		return
	}
	req := new(FindCustomURLReq)
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.Error(err)
		return
	}
	available, err := suh.sus.IsCustomURLAvailable(ctx.Request.Context(), req.CustomCode)
	if err != nil {
		ctx.Error(err)
		return
	}
	if !available {
		ctx.Error(customerror.NewError(
			"url already exist",
			errors.New("custom code already exist"),
			customerror.InvalidAction,
		))
		return
	}
	ctx.JSON(http.StatusOK, dto.ServerResponse[dto.PlainMessageResponse]{
		Success: true,
		Data: dto.PlainMessageResponse{
			Message: "URL available",
		},
	})
}

func (suh *URLHandlerImpl) GetUserLinks(ctx *gin.Context) {
	uid := handlers.GetAuthenticationPayload(ctx)
	if len(uid) == 0 {
		ctx.Error(customerror.NewError(
			"unauthorized",
			errors.New("user id does not provided"),
			customerror.Unauthenticate,
		))
		return
	}
	req := new(GetUserLinksReq)
	if err := ctx.Bind(req); err != nil {
		ctx.Error(err)
		return
	}
	var reqLimit int64 = constant.DEFAULT_LIMIT_PAGINATION
	if req.Limit != nil {
		reqLimit = *req.Limit
	}
	urls, err := suh.sus.GetUserLinks(ctx.Request.Context(), uid, req.LastID, reqLimit)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := []URLRes{}
	for _, u := range urls {
		res = append(res, URLRes(u))
	}
	var lastID *int64 = nil

	if len(res) > 0 {
		lastID = &res[len(res)-1].ID
	}

	ctx.JSON(http.StatusOK, dto.ServerResponse[dto.PaginateRes[URLRes]]{
		Success: true,
		Data: dto.PaginateRes[URLRes]{
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
				LastID: lastID,
				Limit:  reqLimit,
			},
		},
	})
}

func (suh *URLHandlerImpl) NewShortURL(ctx *gin.Context) {
	var userID *string = nil
	uid := handlers.GetAuthenticationPayload(ctx)
	if len(uid) > 0 {
		userID = &uid
	}

	req := new(NewShortURLReq)

	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.Error(err)
		return
	}
	id, eat, err := suh.sus.NewShortURL(ctx.Request.Context(), userID, req.URL, req.Duration, req.CustomCode)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.ServerResponse[NewShortURLRes]{
		Success: true,
		Data: NewShortURLRes{
			URL:         fmt.Sprintf("%s/%s", os.Getenv("AMARY_REDIRECT_DOMAIN"), id),
			OriginalURL: req.URL,
			ExpiredAt:   eat,
		},
	})
}

func (suh *URLHandlerImpl) VisitOriginalURL(ctx *gin.Context) {
	device := ctx.Request.UserAgent()
	id := ctx.Param("id")

	url, err := suh.sus.VisitOriginalURL(ctx.Request.Context(), id, device)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ServerResponse[VisitOriginalUrlRes]{
		Success: true,
		Data: VisitOriginalUrlRes{
			URL: url,
		},
	})
}
