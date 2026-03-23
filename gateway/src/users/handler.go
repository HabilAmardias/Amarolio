package users

import (
	"amarolio-gateway/src/constants"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

type UserServiceItf interface {
	Login() (string, string, error)
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
