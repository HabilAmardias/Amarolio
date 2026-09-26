package users

import (
	"amarolio-auth/src/customerrors"
	"crypto/rand"
	"math/big"
	"strconv"
)

type OTPGen struct{}

func CreateOTPGenerator() *OTPGen {
	return &OTPGen{}
}

func (o *OTPGen) GenerateOTP() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return strconv.Itoa(int(n.Int64())), nil
}
