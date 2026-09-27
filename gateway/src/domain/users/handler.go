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
	Register(email, password string) (string, error)
	Verify(userID string, token string) (string, error)
	PreLogin(email, password string) (string, error)
	Login(userID string, otp string) (string, string, error)
	RefreshAuth(userID string) (string, error)
	ResendVerification(email string) (string, error)
	GetProfile(userID string) (string, error)
	ResendOTP(userID string) (string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) ResendVerification(ctx fiber.Ctx) error {
	body := new(ResendVerificationReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	msg, err := uh.us.ResendVerification(body.Email)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: msg,
		},
	})
}

func (uh *UserHandlerImpl) PreLogin(ctx fiber.Ctx) error {
	body := new(CredentialsReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	token, err := uh.us.PreLogin(body.Email, body.Password)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    token,
		Expires:  time.Now().Add(2 * constants.AUTH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: "Send OTP Success",
		},
	})
}

func (uh *UserHandlerImpl) ResendOTP(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	token, err := uh.us.ResendOTP(claim.Subject)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    token,
		Expires:  time.Now().Add(2 * constants.AUTH_AGE),
		HTTPOnly: true,
		Secure:   secure,
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: "Send OTP Success",
		},
	})
}

func (uh *UserHandlerImpl) Register(ctx fiber.Ctx) error {
	body := new(CredentialsReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	msg, err := uh.us.Register(body.Email, body.Password)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: msg,
		},
	})
}

func (uh *UserHandlerImpl) Verify(ctx fiber.Ctx) error {
	body := new(VerifyReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	msg, err := uh.us.Verify(body.UserID, body.Token)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: msg,
		},
	})
}

func (uh *UserHandlerImpl) LogOut(ctx fiber.Ctx) error {
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
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: "Logged out successfully",
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

func (uh *UserHandlerImpl) Login(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}

	req := new(LoginReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}
	authToken, refreshToken, err := uh.us.Login(claim.Subject, req.OTP)
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
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: "Login successful",
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
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: "refresh token success",
		},
	})
}
