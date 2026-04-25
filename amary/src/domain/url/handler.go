package url

import (
	"amary/src/customerror"
	"amary/src/dto"
	"amary/src/handlers"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type URLServiceItf interface {
	NewShortURL(ctx context.Context, userID *string, longURL string, duration *int) (string, error)
	FindLongURL(ctx context.Context, encodedID string, device string) (string, error)
	GetUserLinks(ctx context.Context, userID string, page int64, limit int64) ([]DecryptedURL, error)
}

type URLHandlerImpl struct {
	sus URLServiceItf
}

func NewURLHandler(sus URLServiceItf) *URLHandlerImpl {
	return &URLHandlerImpl{sus}
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

	urls, err := suh.sus.GetUserLinks(ctx.Request.Context(), uid, req.Page, req.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	res := []UserLinkRes{}
	for _, u := range urls {
		res = append(res, UserLinkRes(u))
	}

	ctx.JSON(http.StatusOK, dto.ServerResponse[[]UserLinkRes]{
		Success: true,
		Data:    res,
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
	id, err := suh.sus.NewShortURL(ctx.Request.Context(), userID, req.URL, req.Duration)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.ServerResponse[NewShortURLRes]{
		Success: true,
		Data: NewShortURLRes{
			URL: fmt.Sprintf("%s/%s", os.Getenv("AMARY_CLIENT_DOMAIN"), id),
		},
	})
}

func (suh *URLHandlerImpl) FindLongURL(ctx *gin.Context) {
	device := ctx.Request.UserAgent()
	id := ctx.Param("id")

	url, err := suh.sus.FindLongURL(ctx.Request.Context(), id, device)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ServerResponse[FindLongUrlRes]{
		Success: true,
		Data: FindLongUrlRes{
			URL: url,
		},
	})
}
