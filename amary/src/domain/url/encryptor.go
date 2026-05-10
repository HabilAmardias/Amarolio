package url

import (
	"amary/src/customerror"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
)

type URLEncryptor struct {
	gcm cipher.AEAD
}

func NewURLEncryptor() *URLEncryptor {
	keyHex := os.Getenv("URL_ENCRYPTION_KEY")
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		log.Fatalln(err.Error())
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &URLEncryptor{gcm}
}

func (ue *URLEncryptor) DecryptURL(cipherURL string) (string, error) {
	data, err := hex.DecodeString(cipherURL)
	if err != nil {
		return "", customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	nonceSize := ue.gcm.NonceSize()
	if len(data) < nonceSize {
		return "", customerror.NewError(
			"something went wrong",
			errors.New("ciphertext too short"),
			customerror.CommonErr,
		)
	}
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := ue.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return string(plaintext), nil
}

func (ue *URLEncryptor) EncryptURL(plainURL string) (string, error) {
	nonce := make([]byte, ue.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	ciphertext := ue.gcm.Seal(nonce, nonce, []byte(plainURL), nil)
	return hex.EncodeToString(ciphertext), nil
}
