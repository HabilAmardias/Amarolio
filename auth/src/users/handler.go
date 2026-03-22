package users

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/dto"
	"amarolio-auth/src/handlers"
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
)

type UserServiceItf interface {
	Login() (string, string)
	RefreshAuth(ctx context.Context, userID string) (string, error)
	LoginCallback(ctx context.Context, code string) (string, string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) Login(ctx fiber.Ctx) error {
	url, state := uh.us.Login()
	return ctx.JSON(dto.ServerResponse{
		Success: true,
		Data: LoginRes{
			URL:   url,
			State: state,
		},
	})
}

func (uh *UserHandlerImpl) RefreshAuth(ctx fiber.Ctx) error {
	userID, err := handlers.GetAuth(ctx)
	if err != nil {
		return err
	}
	authToken, err := uh.us.RefreshAuth(ctx.RequestCtx(), userID)
	if err != nil {
		return err
	}
	return ctx.JSON(dto.ServerResponse{
		Success: true,
		Data: RefreshAuthRes{
			AuthToken: authToken,
		},
	})
}

func (uh *UserHandlerImpl) LoginCallback(ctx fiber.Ctx) error {
	req := new(LoginCallbackReq)
	code := ctx.Query("code")
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}

	if len(code) == 0 {
		return customerrors.NewError(
			"failed to login",
			errors.New("missing code in query param"),
			customerrors.InvalidAction,
		)
	}

	if ctx.Query("state") != req.State {
		return customerrors.NewError(
			"invalid credential",
			errors.New("mismatch state parameters"),
			customerrors.InvalidAction,
		)
	}

	authToken, refreshToken, err := uh.us.LoginCallback(ctx.RequestCtx(), code)
	if err != nil {
		return err
	}
	return ctx.JSON(dto.ServerResponse{
		Success: true,
		Data: LoginCallbackRes{
			AuthToken:    authToken,
			RefreshToken: refreshToken,
		},
	})
}
