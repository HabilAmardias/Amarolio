package users

import (
	"amarolio-auth/src/dto"
	"amarolio-auth/src/handlers"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserServiceItf interface {
	Login(ctx context.Context, challengeID string, otp string) (string, string, error)
	RefreshAuth(ctx context.Context, userID string) (string, error)
	VerifyUser(ctx context.Context, userID, token string) error
	GetProfile(ctx context.Context, userID string) (string, error)
	Register(ctx context.Context, email, password string) error
	ResendOTP(ctx context.Context, challengeID string) (string, error)
	PreLogin(ctx context.Context, email string, password string) (string, error)
	ResendVerification(ctx context.Context, email string) error
	SendResetPasswordEmail(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, userID string, token string, newPassword string) error
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
	if err := uh.us.ResetPassword(ctx.RequestCtx(), body.UserID, body.Token, body.NewPassword); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: dto.TextResponse{
			Message: "Password Reset Successfully",
		},
	})
}

func (uh *UserHandlerImpl) SendResetPasswordEmail(ctx fiber.Ctx) error {
	body := new(SendEmailReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	if err := uh.us.SendResetPasswordEmail(ctx.RequestCtx(), body.Email); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: dto.TextResponse{
			Message: "Reset Password URL Sent",
		},
	})
}

func (uh *UserHandlerImpl) ResendVerification(ctx fiber.Ctx) error {
	body := new(SendEmailReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	if err := uh.us.ResendVerification(ctx.RequestCtx(), body.Email); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: dto.TextResponse{
			Message: "User Verification URL Sent",
		},
	})
}

func (uh *UserHandlerImpl) PreLogin(ctx fiber.Ctx) error {
	body := new(CredentialsReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	token, err := uh.us.PreLogin(ctx.RequestCtx(), body.Email, body.Password)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: OTPRes{
			OTPToken: token,
		},
	})
}

func (uh *UserHandlerImpl) ResendOTP(ctx fiber.Ctx) error {
	userID, err := handlers.GetAuth(ctx)
	if err != nil {
		return err
	}
	token, err := uh.us.ResendOTP(ctx.RequestCtx(), userID)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: OTPRes{
			OTPToken: token,
		},
	})
}

func (uh *UserHandlerImpl) Register(ctx fiber.Ctx) error {
	body := new(CredentialsReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	if err := uh.us.Register(ctx.RequestCtx(), body.Email, body.Password); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: dto.TextResponse{
			Message: "User Register Success",
		},
	})
}

func (uh *UserHandlerImpl) Verify(ctx fiber.Ctx) error {
	body := new(VerifyReq)
	if err := ctx.Bind().JSON(body); err != nil {
		return err
	}
	if err := uh.us.VerifyUser(ctx.RequestCtx(), body.UserID, body.Token); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: dto.TextResponse{
			Message: "Verify User Success",
		},
	})
}

func (uh *UserHandlerImpl) GetProfile(ctx fiber.Ctx) error {
	userID, err := handlers.GetAuth(ctx)
	if err != nil {
		return err
	}

	username, err := uh.us.GetProfile(ctx.RequestCtx(), userID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: GetProfileRes{
			Username: username,
		},
	})
}

func (uh *UserHandlerImpl) Login(ctx fiber.Ctx) error {
	userID, err := handlers.GetAuth(ctx)
	if err != nil {
		return err
	}
	req := new(LoginReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return err
	}

	authToken, refreshToken, err := uh.us.Login(ctx.RequestCtx(), userID, req.OTP)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: LoginRes{
			AuthToken:    authToken,
			RefreshToken: refreshToken,
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
	return ctx.Status(http.StatusOK).JSON(dto.ServerResponse{
		Success: true,
		Data: RefreshAuthRes{
			AuthToken: authToken,
		},
	})
}
