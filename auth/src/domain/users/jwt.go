package users

import (
	"amarolio-auth/src/customerrors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct{}

func CreateJWTUtil() *JWTUtil {
	return &JWTUtil{}
}

type CustomClaim struct {
	jwt.RegisteredClaims
	For int
}

func (ju *JWTUtil) GenerateJWT(id string, usedFor int, age time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claim := CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			Issuer:    os.Getenv("JWT_ISSUER"),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(age)),
		},
		For: usedFor,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", customerrors.NewError(
			"failed to authorize",
			err,
			customerrors.CommonErr,
		)
	}
	return tokenstr, nil
}
