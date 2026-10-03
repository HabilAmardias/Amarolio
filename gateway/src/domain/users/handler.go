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
	Verify(token string) (string, error)
	PreLogin(email, password string) (string, error)
	Login(challengeID string, otp string) (string, string, error)
	RefreshAuth(userID string) (string, error)
	ResendVerification(email string) (string, error)
	GetProfile(userID string) (string, error)
	ResendOTP(challengeID string) (string, error)
	SendResetPasswordEmail(email string) (string, error)
	ResetPassword(userID string, token string, newPassword string) (string, error)
}

type UserHandlerImpl struct {
	us UserServiceItf
}

func NewUserHandler(us UserServiceItf) *UserHandlerImpl {
	return &UserHandlerImpl{us}
}

func (uh *UserHandlerImpl) ResetPassword(ctx fiber.Ctx) error {
	body := new(ResetPasswordReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	msg, err := uh.us.ResetPassword(body.UserID, body.Token, body.NewPassword)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.OTP_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse[dto.TextResponse]{
		Success: true,
		Data: dto.TextResponse{
			Message: msg,
		},
	})
}

func (uh *UserHandlerImpl) SendResetPasswordEmail(ctx fiber.Ctx) error {
	body := new(SendEmailReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	msg, err := uh.us.SendResetPasswordEmail(body.Email)
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

func (uh *UserHandlerImpl) ResendVerification(ctx fiber.Ctx) error {
	body := new(SendEmailReq)
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
	challengeID, err := uh.us.PreLogin(body.Email, body.Password)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.OTP_TOKEN,
		Value:    challengeID,
		Expires:  time.Now().Add(2 * time.Minute),
		HTTPOnly: true,
		SameSite: "Lax",
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
	challengeID := ctx.Cookies(constants.OTP_TOKEN)
	newChallengeID, err := uh.us.ResendOTP(challengeID)
	if err != nil {
		return err
	}
	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.OTP_TOKEN,
		Value:    newChallengeID,
		Expires:  time.Now().Add(2 * time.Minute),
		HTTPOnly: true,
		SameSite: "Lax",
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
	msg, err := uh.us.Verify(body.Token)
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
		SameSite: "Lax",
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
		Expires:  time.Now().Add(-3 * time.Minute),
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.OTP_TOKEN,
		Value:    "",
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "Lax",
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
	challengeID := ctx.Cookies(constants.OTP_TOKEN)
	req := new(LoginReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}
	authToken, refreshToken, err := uh.us.Login(challengeID, req.OTP)
	if err != nil {
		return err
	}

	secure := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.AUTH_TOKEN,
		Value:    authToken,
		Expires:  time.Now().Add(2 * constants.AUTH_AGE),
		HTTPOnly: true,
		SameSite: "Lax",
		Secure:   secure,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.REFRESH_TOKEN,
		Value:    refreshToken,
		Expires:  time.Now().Add(2 * constants.REFRESH_AGE),
		HTTPOnly: true,
		SameSite: "Lax",
		Secure:   secure,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     constants.OTP_TOKEN,
		Value:    "",
		HTTPOnly: true,
		SameSite: "Lax",
		Secure:   secure,
		Expires:  time.Now().Add(-3 * time.Minute),
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
		SameSite: "Lax",
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
