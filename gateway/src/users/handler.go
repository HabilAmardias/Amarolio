package users

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/utils"
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

func getAuthPayload(ctx fiber.Ctx, key string) (*utils.CustomClaim, error) {
	claim, ok := ctx.Locals(key).(*utils.CustomClaim)
	if !ok {
		return nil, customerrors.NewError(
			"credentials does not found",
			errors.New("jwt claim is missing"),
			customerrors.Unauthenticate,
		)
	}
	return claim, nil
}

type UserServiceItf interface {
	Login() (string, string, error)
	RefreshAuth(userID string) (string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) Login(ctx fiber.Ctx) error {
	var isProd bool = os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	state, url, err := uh.us.Login()
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		MaxAge:   int(30 * time.Second),
		Name:     "oauthstate",
		Value:    state,
		HTTPOnly: true,
		Secure:   isProd,
	})
	return ctx.Redirect().To(url)
}

func (uh *UserHandlerImpl) RefreshAuth(ctx fiber.Ctx) error {
	claim, err := getAuthPayload(ctx, constants.REFRESH_CLAIM_KEY)
	if err != nil {
		return err
	}
	authToken, err := uh.us.RefreshAuth(claim.Subject)
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    authToken,
		HTTPOnly: true,
		MaxAge:   int(constants.AUTH_AGE),
		Secure:   os.Getenv("ENVIRONMENT") == constants.PRODUCTION,
	})
	return ctx.JSON(dto.ServerResponse[RefreshAuthRes]{
		Success: true,
		Data: RefreshAuthRes{
			Message: "refresh token success",
		},
	})
}
