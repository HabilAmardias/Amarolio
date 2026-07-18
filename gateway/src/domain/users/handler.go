package users

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type UserServiceItf interface {
	Login() (string, string, error)
	RefreshAuth(userID string) (string, error)
	LoginCallback(code string, state string) (string, string, error)
	GetProfile(userID string) (string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) LogOut(ctx fiber.Ctx) error {
	req := new(LogoutReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[LogoutRes]{
		Success: true,
		Data: LogoutRes{
			RedirectURI: req.RedirectURI,
		},
	})
}

func (uh *UserHandlerImpl) GetProfile(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	username, err := uh.us.GetProfile(claim.Subject)
	if err != nil {
		return err
	}
	return ctx.Status(fasthttp.StatusOK).JSON(dto.ServerResponse[GetProfileRes]{
		Success: true,
		Data: GetProfileRes{
			Username: username,
		},
	})
}

func (uh *UserHandlerImpl) LoginCallback(ctx fiber.Ctx) error {
	state := ctx.Cookies("oauthstate")
	code := ctx.Query("code")
	redirectURI := ctx.Cookies("redirect_uri")

	authToken, refreshToken, err := uh.us.LoginCallback(code, state)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    authToken,
		Expires:  time.Now().Add(2 * constants.AUTH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    refreshToken,
		Expires:  time.Now().Add(constants.REFRESH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})

	// remove redirect_uri and oauthstate cookies
	ctx.Cookie(&fiber.Cookie{
		Name:     "redirect_uri",
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	return ctx.Redirect().Status(http.StatusTemporaryRedirect).To(redirectURI)
}

func (uh *UserHandlerImpl) Login(ctx fiber.Ctx) error {
	req := new(LoginReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}
	var isProd bool = os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	state, url, err := uh.us.Login()
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "redirect_uri",
		Expires:  time.Now().Add(constants.AUTH_AGE),
		Value:    req.RedirectURI,
		HTTPOnly: true,
		Secure:   isProd,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Expires:  time.Now().Add(30 * time.Second),
		Value:    state,
		HTTPOnly: true,
		Secure:   isProd,
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[LoginRes]{
		Success: true,
		Data: LoginRes{
			RedirectURI: url,
		},
	})
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
		Expires:  time.Now().Add(constants.AUTH_AGE),
		Secure:   os.Getenv("ENVIRONMENT") == constants.PRODUCTION,
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[RefreshAuthRes]{
		Success: true,
		Data: RefreshAuthRes{
			Message: "refresh token success",
		},
	})
}
