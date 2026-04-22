package users

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

type UserServiceItf interface {
	Login() (string, string, error)
	RefreshAuth(userID string) (string, error)
	LoginCallback(code string, state string) (string, string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) LoginCallback(ctx fiber.Ctx) error {
	state := ctx.Cookies("oauthstate")
	code := ctx.Query("code")
	clientDomain, exist := os.LookupEnv("CLIENT_DOMAIN")
	if !exist {
		return customerrors.NewError(
			"something went wrong",
			errors.New("client domain does not exist"),
			customerrors.CommonErr,
		)
	}

	authToken, refreshToken, err := uh.us.LoginCallback(code, state)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    authToken,
		MaxAge:   int(constants.AUTH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    refreshToken,
		MaxAge:   int(constants.REFRESH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})
	return ctx.Redirect().To(clientDomain + "/login-callback")
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
	return ctx.Status(http.StatusPermanentRedirect).Redirect().To(url)
}

func (uh *UserHandlerImpl) RefreshAuth(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.REFRESH_CLAIM_KEY)
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
