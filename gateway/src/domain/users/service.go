package users

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/services"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type UserServiceImpl struct {
	hs string
	pr string
}

func NewUserService(hs string, pr string) *UserServiceImpl {
	return &UserServiceImpl{hs, pr}
}

func (us *UserServiceImpl) callLogin(userID string, otp string) (*dto.ServerResponse[Login], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	reqBody := LoginBody{
		OTP: otp,
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[Login](us.hs, us.pr, "/api/v1/login", fasthttp.MethodPost, fasthttp.StatusOK, b, nil, headers)
}

func (us *UserServiceImpl) callRefreshAuth(userID string) (*dto.ServerResponse[RefreshAuth], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	return services.Call[RefreshAuth](us.hs, us.pr, "/api/v1/refresh", fasthttp.MethodPost, fasthttp.StatusOK, nil, nil, headers)
}

func (us *UserServiceImpl) callGetProfile(userID string) (*dto.ServerResponse[GetProfile], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	return services.Call[GetProfile](us.hs, us.pr, "/api/v1/me", fasthttp.MethodGet, fasthttp.StatusOK, nil, nil, headers)
}

func (us *UserServiceImpl) callVerify(userID string, token string) (*dto.ServerResponse[Text], error) {
	body := VerificationBody{
		UserID: userID,
		Token:  token,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[Text](us.hs, us.pr, "/api/v1/verify", fasthttp.MethodPost, fasthttp.StatusOK, b, nil, nil)
}

func (us *UserServiceImpl) callRegister(email, password string) (*dto.ServerResponse[Text], error) {
	body := CredentialsBody{
		Email:    email,
		Password: password,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[Text](us.hs, us.pr, "/api/v1/register", fasthttp.MethodPost, fasthttp.StatusOK, b, nil, nil)
}

func (us *UserServiceImpl) callResendOTP(userID string) (*dto.ServerResponse[OTP], error) {
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	return services.Call[OTP](us.hs, us.pr, "/api/v1/otp/send", fasthttp.MethodGet, fasthttp.StatusOK, nil, nil, headers)
}

func (us *UserServiceImpl) callPreLogin(email, password string) (*dto.ServerResponse[OTP], error) {
	body := CredentialsBody{
		Email:    email,
		Password: password,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[OTP](us.hs, us.pr, "/api/v1/prelogin", fasthttp.MethodPost, fasthttp.StatusOK, b, nil, nil)
}

func (us *UserServiceImpl) callResendVerification(email string) (*dto.ServerResponse[Text], error) {
	body := ResendVerificationBody{
		Email: email,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return services.Call[Text](us.hs, us.pr, "/api/v1/verify/send", fasthttp.MethodPost, fasthttp.StatusOK, b, nil, nil)
}

func (us *UserServiceImpl) ResendVerification(email string) (string, error) {
	res, err := us.callResendVerification(email)
	if err != nil {
		return "", err
	}
	return res.Data.Message, nil
}

func (us *UserServiceImpl) PreLogin(email, password string) (string, error) {
	res, err := us.callPreLogin(email, password)
	if err != nil {
		return "", err
	}
	return res.Data.OTPToken, nil
}

func (us *UserServiceImpl) ResendOTP(userID string) (string, error) {
	res, err := us.callResendOTP(userID)
	if err != nil {
		return "", err
	}
	return res.Data.OTPToken, nil
}

func (us *UserServiceImpl) Register(email, password string) (string, error) {
	res, err := us.callRegister(email, password)
	if err != nil {
		return "", err
	}
	return res.Data.Message, nil
}

func (us *UserServiceImpl) Verify(userID string, token string) (string, error) {
	res, err := us.callVerify(userID, token)
	if err != nil {
		return "", err
	}
	return res.Data.Message, nil
}

func (us *UserServiceImpl) GetProfile(userID string) (string, error) {
	res, err := us.callGetProfile(userID)
	if err != nil {
		return "", err
	}
	return res.Data.Username, nil
}

func (us *UserServiceImpl) Login(userID string, otp string) (string, string, error) {
	res, err := us.callLogin(userID, otp)
	if err != nil {
		return "", "", err
	}
	return res.Data.AuthToken, res.Data.RefreshToken, nil
}

func (us *UserServiceImpl) RefreshAuth(userID string) (string, error) {
	res, err := us.callRefreshAuth(userID)
	if err != nil {
		return "", err
	}

	return res.Data.Token, nil
}
