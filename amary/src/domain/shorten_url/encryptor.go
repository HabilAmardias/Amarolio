package shortenurl

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

type URLEncryptor struct {
	gcm cipher.AEAD
}

func NewURLEncryptor() *URLEncryptor {
	key := os.Getenv("URL_ENCRYPTION_KEY")
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
	nonce := make([]byte, ue.gcm.NonceSize())
	plainText, err := ue.gcm.Open(nonce, nonce, []byte(cipherURL), nil)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func (ue *URLEncryptor) EncryptURL(plainURL string) string {
	nonce := make([]byte, ue.gcm.NonceSize())
	ciphertext := ue.gcm.Seal(nonce, nonce, []byte(plainURL), nil)

	return string(ciphertext)
}
