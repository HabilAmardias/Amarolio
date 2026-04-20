package shortenurl

import (
	"amary/src/dto"
	"amary/src/handlers"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ShortenURLServiceItf interface {
	NewShortURL(ctx context.Context, userID *string, longURL string, duration *int) (string, error)
	FindLongURL(ctx context.Context, encodedID string, device string) (string, error)
}

type ShortenURLHandlerImpl struct {
	sus ShortenURLServiceItf
}

func NewShortenURLHandler(sus ShortenURLServiceItf) *ShortenURLHandlerImpl {
	return &ShortenURLHandlerImpl{sus}
}

func (suh *ShortenURLHandlerImpl) NewShortURL(ctx *gin.Context) {
	userID := new(string)
	req := new(NewShortURLReq)
	uid := handlers.GetAuthenticationPayload(ctx)
	if len(uid) == 0 {
		userID = nil
	}
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

func (suh *ShortenURLHandlerImpl) FindLongURL(ctx *gin.Context) {
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
