package users

import (
	"amarolio-auth/src/customerrors"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type Hasher struct{}

func CreateHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) HashPassword(password string) (string, error) {
	return h.generateFromPassword(password)
}

func (h *Hasher) ValidatePassword(encodedHash, password string) (bool, error) {
	p, salt, hash, err := h.decodeHash(encodedHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (h *Hasher) generateFromPassword(password string) (encodedHash string, err error) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	salt, err := h.generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (h *Hasher) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, customerrors.NewError(
			"Something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	return b, nil
}
func (h *Hasher) decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, customerrors.NewError(
			"Something went wrong",
			errors.New("invalid hash"),
			customerrors.CommonErr,
		)
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	if version != argon2.Version {
		return nil, nil, nil, customerrors.NewError(
			"Something went wrong",
			errors.New("invalid argon2 version"),
			customerrors.CommonErr,
		)
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
