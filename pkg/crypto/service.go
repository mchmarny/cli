package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/pkg/errors"
)

var (
	bytes = []byte{23, 27, 88, 76, 45, 13, 73, 25, 83, 78, 53, 62, 7, 57, 71, 20}
)

// New creates a new Crypto service instance.
func New(key string) (*Service, error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher block")
	}

	return &Service{
		block: block,
	}, nil
}

// Service provides twitter service.
type Service struct {
	block cipher.Block
}

// Encrypt encrypts the provided value using service-level key
func (s *Service) Encrypt(val string) (string, error) {
	clearText := []byte(val)
	cfb := cipher.NewCFBEncrypter(s.block, bytes)
	cipherText := make([]byte, len(clearText))
	cfb.XORKeyStream(cipherText, clearText)
	return encode(cipherText), nil
}

// Decrypt decrypts previously encrypted value.
func (s *Service) Decrypt(val string) (string, error) {
	cipherText, err := decode(val)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode cipher text")
	}
	cfb := cipher.NewCFBDecrypter(s.block, bytes)
	clearText := make([]byte, len(cipherText))
	cfb.XORKeyStream(clearText, cipherText)
	return string(clearText), nil
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 string")
	}
	return data, nil
}
