package utils

import (
	"amarolio-gateway/src/customerrors"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct{}

type CustomClaim struct {
	jwt.RegisteredClaims
	For int
}

func NewJWTUtil() *JWTUtil {
	return &JWTUtil{}
}

func (ju *JWTUtil) VerifyJWT(tokenStr string, usedFor int) (*CustomClaim, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	custClaim := new(CustomClaim)

	token, err := jwt.ParseWithClaims(tokenStr, custClaim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, customerrors.NewError(
				"Failed to authorize",
				fmt.Errorf("failed to parse jwt"),
				customerrors.Unauthenticate,
			)
		}
		return []byte(jwtSecret), nil
	},
		jwt.WithIssuer(os.Getenv("JWT_ISSUER")),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, customerrors.NewError(
				"session expired",
				err,
				customerrors.Unauthenticate,
			)
		}
		return nil, customerrors.NewError(
			"Failed to authorize",
			err,
			customerrors.Unauthenticate,
		)
	}
	if !token.Valid {
		return nil, customerrors.NewError(
			"Failed to authorize",
			fmt.Errorf("jwt token is not valid"),
			customerrors.Unauthenticate,
		)
	}
	claim, ok := token.Claims.(*CustomClaim)
	if !ok {
		return nil, customerrors.NewError(
			"Failed to authorize",
			fmt.Errorf("failed to get jwt claim"),
			customerrors.Unauthenticate,
		)
	}
	if claim.For != usedFor {
		return nil, customerrors.NewError(
			"invalid token",
			errors.New("invalid token usage"),
			customerrors.InvalidAction,
		)
	}
	return claim, nil
}
