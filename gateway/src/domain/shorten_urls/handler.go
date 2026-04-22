package shortenurls

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type ShortenURLServiceItf interface {
	FindLongURL(id string, device string) (string, error)
	NewShortURL(userID *string, url string, duration *int) (string, error)
}

type ShortenURLHandlerImpl struct {
	sus ShortenURLServiceItf
}

func NewShortenURLHandler(sus ShortenURLServiceItf) *ShortenURLHandlerImpl {
	return &ShortenURLHandlerImpl{sus}
}

func (suh *ShortenURLHandlerImpl) RedirectToURL(ctx fiber.Ctx) error {
	device := ctx.Request().Header.UserAgent()
	id := ctx.Params("id")

	url, err := suh.sus.FindLongURL(id, string(device))
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusFound).Redirect().To(url)
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

	url, err := suh.sus.NewShortURL(userID, req.URL, req.Duration)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(dto.ServerResponse[NewShortenURLRes]{
		Success: true,
		Data: NewShortenURLRes{
			URL: url,
		},
	})
}
